## Installation

### Kubernetes

```sh
kubectl apply -f https://raw.githubusercontent.com/drorIvry/rego/main/deploy/deployment.yml
```

### Helm

```sh
helm repo add rego https://drorivry.github.io/rego-charts/
helm install --generate-name rego/rego
```

### From source code

1. clone the [repo](https://drorivry.github.io/rego) 
2. build using go
    ```sh
    go build ./
    ```
3. add execution permissions
    ```sh
    sudo chmod +x rego
    ```
4. create a `.env` file with the following
   ```sh
   PORT=4004
    DB_URL=rego.db
    TASK_TIMEOUT=300
    IN_CLUSTER=false
    ```
### Installing the rego CLI (Recommended)

```sh
curl -L https://raw.githubusercontent.com/drorIvry/rego-cli/main/install.sh | sh
```

## Core concepts

### Task Definitions 

a task definition is the logical representation of a task, it is not the actual run of a specific task but the configuration of the task.

A task definition can be periodic or a one time. it can contain metadata and more, 
here is  an example of a task definition

```go
type TaskDefinition struct {
	gorm.Model
	ID                      uuid.UUID      `json:"id" gorm:"type:uuid` //;default:uuid_generate_v4()"
	Image                   string         `json:"image" binding:"required"`
	Name                    string         `json:"name"`
	Namespace               string         `json:"namespace"`
	ExecutionInterval       int            `json:"execution_interval"`
	ExecutionsCounter       int            `json:"execution_counter"`
	NextExecutionTime       time.Time      `json:"next_execution_time"`
	Enabled                 bool           `json:"enabled"`
	Deleted                 bool           `json:"deleted"`
	Cmd                     pq.StringArray `json:"cmd" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata"`
}
```

- `image` is the docker image to run, it can be public image off of docker.io or any private registry as long as it is accessible to your k8s custer
- `execution_interval` - is a marker that if set tells rego to rerun the task every X seconds
- `execution_counter` - how many times did the task run
- `metadata` - this field serves 2 purposes
  1. an open field for users to store task definition specific metadata for example: 
        lets say I want to store the client_id of the client that the task is related to.
  2. A way to configure the k8s job spec with `batchv1.JobSpec` like:

```go
type JobSpec struct {
	Parallelism *int32 `json:"parallelism,omitempty" protobuf:"varint,1,opt,name=parallelism"`
	Completions *int32 `json:"completions,omitempty" protobuf:"varint,2,opt,name=completions"`
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,3,opt,name=activeDeadlineSeconds"`
	PodFailurePolicy *PodFailurePolicy `json:"podFailurePolicy,omitempty" protobuf:"bytes,11,opt,name=podFailurePolicy"`
	BackoffLimit *int32 `json:"backoffLimit,omitempty" protobuf:"varint,7,opt,name=backoffLimit"`
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`
	ManualSelector *bool `json:"manualSelector,omitempty" protobuf:"varint,5,opt,name=manualSelector"`
	Template corev1.PodTemplateSpec `json:"template" protobuf:"bytes,6,opt,name=template"`
	TTLSecondsAfterFinished *int32 `json:"ttlSecondsAfterFinished,omitempty" protobuf:"varint,8,opt,name=ttlSecondsAfterFinished"`
	CompletionMode *CompletionMode `json:"completionMode,omitempty" protobuf:"bytes,9,opt,name=completionMode,casttype=CompletionMode"`
	Suspend *bool `json:"suspend,omitempty" protobuf:"varint,10,opt,name=suspend"`
}
```

### Task Execution

Task execution is the  object representing a specific task run. 
You can think of it like the log in the run record.

```go
type TaskExecution struct {
	gorm.Model
	ID                      uuid.UUID      `json:"id,omitempty" gorm:"type:uuid`
	TaskDefinitionId        uuid.UUID      `json:"task_definition_id,omitempty"`
	StatusCode              Status         `json:"status_code,omitempty"`
	TaskStatus              string         `json:"status,omitempty"`
	Image                   string         `json:"image,omitempty" binding:"required"`
	Name                    string         `json:"name,omitempty"`
	Namespace               string         `json:"namespace,omitempty"`
	Cmd                     pq.StringArray `json:"cmd,omitempty" gorm:"type:text[]"`
	Metadata                datatypes.JSON `json:"metadata,omitempty"`
}
```

many of the fields are shared with the task definition but here are some specific ones.

- `status_code` -  enumeration of task state. 
  - ```go
  const (
	READY        Status = 100
	JOB_DEPLOYED Status = 200
	PENDING      Status = 250
	RUNNING      Status = 300
	TIMEOUT      Status = 400
	PROC_ERROR   Status = 410
	APP_ERROR    Status = 420
	ABORTED      Status = 430
	SUCCESS      Status = 500
)
```
