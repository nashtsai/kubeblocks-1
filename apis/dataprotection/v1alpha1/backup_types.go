/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupSpec defines the desired state of Backup.
type BackupSpec struct {
	// Which backupPolicy is applied to perform this backup
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`
	BackupPolicyName string `json:"backupPolicyName"`

	// Backup Type. datafile or logfile or snapshot. If not set, datafile is the default type.
	// +kubebuilder:default=datafile
	BackupType BackupType `json:"backupType"`

	// if backupType is incremental, parentBackupName is required.
	// +optional
	ParentBackupName string `json:"parentBackupName,omitempty"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	// +optional
	Phase BackupPhase `json:"phase,omitempty"`

	// Records parentBackupName if backupType is incremental.
	// +optional
	ParentBackupName string `json:"parentBackupName,omitempty"`

	// The date and time when the Backup is eligible for garbage collection.
	// 'null' means the Backup is NOT be cleaned except delete manual.
	// +optional
	Expiration *metav1.Time `json:"expiration,omitempty"`

	// Date/time when the backup started being processed.
	// +optional
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// Date/time when the backup finished being processed.
	// +optional
	CompletionTimestamp *metav1.Time `json:"completionTimestamp,omitempty"`

	// The duration time of backup execution.
	// When converted to a string, the form is "1h2m0.5s".
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`

	// Backup total size.
	// A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	// +optional
	TotalSize string `json:"totalSize,omitempty"`

	// The reason for a backup failure.
	// +optional
	FailureReason string `json:"failureReason,omitempty"`

	// remoteVolume saves the backup data.
	// +optional
	PersistentVolumeClaimName string `json:"persistentVolumeClaimName,omitempty"`

	// logFilePersistentVolumeClaimName saves the logfile backup data.
	// +optional
	LogFilePersistentVolumeClaimName string `json:"logFilePersistentVolumeClaimName,omitempty"`

	// backupToolName references the backup tool name.
	// +optional
	BackupToolName string `json:"backupToolName,omitempty"`

	// sourceCluster records the source cluster information for this backup.
	SourceCluster string `json:"sourceCluster,omitempty"`

	// availableReplicas available replicas for statefulSet which created by backup.
	// +optional
	AvailableReplicas *int32 `json:"availableReplicas,omitempty"`

	// manifests determines the backup metadata info.
	// +optional
	Manifests *ManifestsStatus `json:"manifests,omitempty"`
}

type ManifestsStatus struct {
	// backupLog records startTime and stopTime of data logging.
	// +optional
	BackupLog *BackupLogStatus `json:"backupLog,omitempty"`

	// target records the target cluster metadata string, which is in JSON format.
	// +optional
	Target string `json:"target,omitempty"`

	// snapshot records the volume snapshot metadata.
	// +optional
	Snapshot *BackupSnapshotStatus `json:"backupSnapshot,omitempty"`

	// backupTool records information about backup files generated by the backup tool.
	// +optional
	BackupTool *BackupToolManifestsStatus `json:"backupTool,omitempty"`

	// userContext stores some loosely structured and extensible information.
	// +optional
	UserContext map[string]string `json:"userContext,omitempty"`
}

type BackupLogStatus struct {
	// startTime records the start time of data logging.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// stopTime records the stop time of data logging.
	// +optional
	StopTime *metav1.Time `json:"stopTime,omitempty"`
}

type BackupSnapshotStatus struct {
	// volumeSnapshotName records the volumeSnapshot name.
	// +optional
	VolumeSnapshotName string `json:"volumeSnapshotName,omitempty"`

	// volumeSnapshotContentName specifies the name of a pre-existing VolumeSnapshotContent
	// object representing an existing volume snapshot.
	// This field should be set if the snapshot already exists and only needs a representation in Kubernetes.
	// This field is immutable.
	// +optional
	VolumeSnapshotContentName string `json:"volumeSnapshotContentName,omitempty"`
}

type BackupToolManifestsStatus struct {
	// filePath records the file path of backup.
	// +optional
	FilePath string `json:"filePath,omitempty"`

	// logFilePath records the log file path of backup.
	// +optional
	LogFilePath string `json:"logFilePath,omitempty"`

	// volumeName records volume name of backup data pvc.
	// +optional
	VolumeName string `json:"volumeName,omitempty"`

	// Backup upload total size.
	// A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	// +optional
	UploadTotalSize string `json:"uploadTotalSize,omitempty"`

	// checksum of backup file, generated by md5 or sha1 or sha256.
	// +optional
	Checksum string `json:"checksum,omitempty"`

	// backup checkpoint, for incremental backup.
	// +optional
	Checkpoint string `json:"checkpoint,omitempty"`
}

// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories={kubeblocks},scope=Namespaced
// +kubebuilder:printcolumn:name="TYPE",type=string,JSONPath=`.spec.backupType`
// +kubebuilder:printcolumn:name="STATUS",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="SOURCE-CLUSTER",type=string,JSONPath=`.status.sourceCluster`
// +kubebuilder:printcolumn:name="TOTAL-SIZE",type=string,JSONPath=`.status.totalSize`
// +kubebuilder:printcolumn:name="DURATION",type=string,JSONPath=`.status.duration`
// +kubebuilder:printcolumn:name="CREATE-TIME",type=string,JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="COMPLETION-TIME",type=string,JSONPath=`.status.completionTimestamp`

// Backup is the Schema for the backups API (defined by User).
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backup.
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}

// Validate validates the BackupSpec and returns an error if invalid.
func (r *BackupSpec) Validate(backupPolicy *BackupPolicy) error {
	notSupportedMessage := "backupPolicy: %s not supports %s backup in backupPolicy"
	switch r.BackupType {
	case BackupTypeSnapshot:
		if backupPolicy.Spec.Snapshot == nil {
			return fmt.Errorf(notSupportedMessage, r.BackupPolicyName, BackupTypeSnapshot)
		}
	case BackupTypeDataFile:
		if backupPolicy.Spec.Datafile == nil {
			return fmt.Errorf(notSupportedMessage, r.BackupPolicyName, BackupTypeDataFile)
		}
	case BackupTypeLogFile:
		if backupPolicy.Spec.Logfile == nil {
			return fmt.Errorf(notSupportedMessage, r.BackupPolicyName, BackupTypeLogFile)
		}
	}
	return nil
}

// GetRecoverableTimeRange returns the recoverable time range array.
func GetRecoverableTimeRange(backups []Backup) []BackupLogStatus {
	sort.Slice(backups, func(i, j int) bool {
		if backups[i].Status.StartTimestamp == nil && backups[j].Status.StartTimestamp != nil {
			return false
		}
		if backups[i].Status.StartTimestamp != nil && backups[j].Status.StartTimestamp == nil {
			return true
		}
		if backups[i].Status.StartTimestamp.Equal(backups[j].Status.StartTimestamp) {
			return backups[i].Name < backups[j].Name
		}
		return backups[i].Status.StartTimestamp.Before(backups[j].Status.StartTimestamp)
	})
	getLogfileStartTimeAndStopTime := func() (metav1.Time, metav1.Time) {
		var (
			startTime metav1.Time
			stopTime  metav1.Time
		)
		for _, b := range backups {
			if b.Status.Manifests == nil || b.Status.Manifests.BackupLog == nil ||
				b.Status.Manifests.BackupLog.StopTime == nil ||
				b.Status.Manifests.BackupLog.StartTime == nil {
				continue
			}
			if b.Spec.BackupType != BackupTypeLogFile {
				continue
			}
			if startTime.IsZero() {
				startTime = *b.Status.Manifests.BackupLog.StartTime
			}
			stopTime = *b.Status.Manifests.BackupLog.StopTime
		}
		return startTime, stopTime
	}
	logfileStartTime, logfileStopTime := getLogfileStartTimeAndStopTime()
	// if not exists the startTime/stopTime of the first log file, return
	if logfileStartTime.IsZero() || logfileStopTime.IsZero() {
		return nil
	}
	getFirstRecoverableBaseBackup := func() *Backup {
		for _, b := range backups {
			if !slices.Contains([]BackupType{BackupTypeDataFile, BackupTypeSnapshot}, b.Spec.BackupType) ||
				b.Status.Phase != BackupCompleted {
				continue
			}
			if b.Status.Manifests == nil || b.Status.Manifests.BackupLog == nil ||
				b.Status.Manifests.BackupLog.StopTime == nil {
				continue
			}
			// checks if the baseBackup stop time is between logfileStartTime and logfileStopTime.
			if !b.Status.Manifests.BackupLog.StopTime.Before(&logfileStartTime) &&
				b.Status.Manifests.BackupLog.StopTime.Before(&logfileStopTime) {
				return &b
			}
		}
		return nil
	}
	firstRecoverableBaseBackup := getFirstRecoverableBaseBackup()
	if firstRecoverableBaseBackup == nil {
		return nil
	}
	// range of recoverable time
	return []BackupLogStatus{{StopTime: &logfileStopTime,
		StartTime: firstRecoverableBaseBackup.Status.Manifests.BackupLog.StopTime}}
}
