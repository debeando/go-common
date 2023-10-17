package mongodb

import "time"

type Databases struct {
	Databases     []Database `bson:"databases"`
	TotalSize     int64      `bson:"totalSize"`
	Ok            int64      `bson:"ok"`
	OperationTime time.Time  `bson:"operationTime"`
}

type Database struct {
	Name       string `bson:"name"`
	SizeOnDisk int64  `bson:"sizeOnDisk"`
	Empty      bool   `bson:"empty"`
}

type ServerStatus struct {
	SampleTime         time.Time              `bson:""`
	Flattened          map[string]interface{} `bson:""`
	Host               string                 `bson:"host"`
	Version            string                 `bson:"version"`
	Process            string                 `bson:"process"`
	Pid                int64                  `bson:"pid"`
	Uptime             int64                  `bson:"uptime"`
	UptimeMillis       int64                  `bson:"uptimeMillis"`
	UptimeEstimate     int64                  `bson:"uptimeEstimate"`
	LocalTime          time.Time              `bson:"localTime"`
	Asserts            *AssertsStats          `bson:"asserts"`
	BackgroundFlushing *FlushStats            `bson:"backgroundFlushing"`
	ExtraInfo          *ExtraInfo             `bson:"extra_info"`
	Connections        *ConnectionStats       `bson:"connections"`
	Dur                *DurStats              `bson:"dur"`
	GlobalLock         *GlobalLockStats       `bson:"globalLock"`
	Locks              map[string]LockStats   `bson:"locks,omitempty"`
	Network            *NetworkStats          `bson:"network"`
	Opcounters         *OpcountStats          `bson:"opcounters"`
	OpcountersRepl     *OpcountStats          `bson:"opcountersRepl"`
	OpLatencies        *OpLatenciesStats      `bson:"opLatencies"`
	RecordStats        *DBRecordStats         `bson:"recordStats"`
	Mem                *MemStats              `bson:"mem"`
	Repl               *ReplStatus            `bson:"repl"`
	ShardCursorType    map[string]interface{} `bson:"shardCursorType"`
	StorageEngine      *StorageEngine         `bson:"storageEngine"`
	WiredTiger         *WiredTiger            `bson:"wiredTiger"`
	Metrics            *MetricsStats          `bson:"metrics"`
	TCMallocStats      *TCMallocStats         `bson:"tcmalloc"`
}

// AssertsStats stores information related to assertions raised since the MongoDB process started
type AssertsStats struct {
	Regular   int64 `bson:"regular"`
	Warning   int64 `bson:"warning"`
	Msg       int64 `bson:"msg"`
	User      int64 `bson:"user"`
	Rollovers int64 `bson:"rollovers"`
}

// FlushStats stores information about memory flushes.
type FlushStats struct {
	Flushes      int64     `bson:"flushes"`
	TotalMs      int64     `bson:"total_ms"`
	AverageMs    float64   `bson:"average_ms"`
	LastMs       int64     `bson:"last_ms"`
	LastFinished time.Time `bson:"last_finished"`
}

// ExtraInfo stores additional platform specific information.
type ExtraInfo struct {
	PageFaults *int64 `bson:"page_faults"`
}

// ConnectionStats stores information related to incoming database connections.
type ConnectionStats struct {
	Current      int64 `bson:"current"`
	Available    int64 `bson:"available"`
	TotalCreated int64 `bson:"totalCreated"`
}

// DurStats stores information related to journaling statistics.
type DurStats struct {
	Commits            float64 `bson:"commits"`
	JournaledMB        float64 `bson:"journaledMB"`
	WriteToDataFilesMB float64 `bson:"writeToDataFilesMB"`
	Compression        float64 `bson:"compression"`
	CommitsInWriteLock float64 `bson:"commitsInWriteLock"`
	EarlyCommits       float64 `bson:"earlyCommits"`
	TimeMs             DurTiming
}

// DurTiming stores information related to journaling.
type DurTiming struct {
	Dt               int64 `bson:"dt"`
	PrepLogBuffer    int64 `bson:"prepLogBuffer"`
	WriteToJournal   int64 `bson:"writeToJournal"`
	WriteToDataFiles int64 `bson:"writeToDataFiles"`
	RemapPrivateView int64 `bson:"remapPrivateView"`
}

// GlobalLockStats stores information related locks in the MMAP storage engine.
type GlobalLockStats struct {
	TotalTime     int64        `bson:"totalTime"`
	LockTime      int64        `bson:"lockTime"`
	CurrentQueue  *QueueStats  `bson:"currentQueue"`
	ActiveClients *ClientStats `bson:"activeClients"`
}

// QueueStats stores the number of queued read/write operations.
type QueueStats struct {
	Total   int64 `bson:"total"`
	Readers int64 `bson:"readers"`
	Writers int64 `bson:"writers"`
}

// ClientStats stores the number of active read/write operations.
type ClientStats struct {
	Total   int64 `bson:"total"`
	Readers int64 `bson:"readers"`
	Writers int64 `bson:"writers"`
}

// LockStats stores information related to time spent acquiring/holding locks
// for a given database.
type LockStats struct {
	TimeLockedMicros    ReadWriteLockTimes `bson:"timeLockedMicros"`
	TimeAcquiringMicros ReadWriteLockTimes `bson:"timeAcquiringMicros"`

	// AcquireCount and AcquireWaitCount are new fields of the lock stats only populated on 3.0 or newer.
	// Typed as a pointer so that if it is nil, mongostat can assume the field is not populated
	// with real namespace data.
	AcquireCount     *ReadWriteLockTimes `bson:"acquireCount,omitempty"`
	AcquireWaitCount *ReadWriteLockTimes `bson:"acquireWaitCount,omitempty"`
}

// ReadWriteLockTimes stores time spent holding read/write locks.
type ReadWriteLockTimes struct {
	Read       int64 `bson:"R"`
	Write      int64 `bson:"W"`
	ReadLower  int64 `bson:"r"`
	WriteLower int64 `bson:"w"`
}

// NetworkStats stores information related to network traffic.
type NetworkStats struct {
	BytesIn     int64 `bson:"bytesIn"`
	BytesOut    int64 `bson:"bytesOut"`
	NumRequests int64 `bson:"numRequests"`
}

// OpcountStats stores information related to commands and basic CRUD operations.
type OpcountStats struct {
	Insert  int64 `bson:"insert"`
	Query   int64 `bson:"query"`
	Update  int64 `bson:"update"`
	Delete  int64 `bson:"delete"`
	GetMore int64 `bson:"getmore"`
	Command int64 `bson:"command"`
}

// OpLatenciesStats stores information related to operation latencies for the database as a whole
type OpLatenciesStats struct {
	Reads    *LatencyStats `bson:"reads"`
	Writes   *LatencyStats `bson:"writes"`
	Commands *LatencyStats `bson:"commands"`
}

// LatencyStats lists total latency in microseconds and count of operations, enabling you to obtain an average
type LatencyStats struct {
	Latency int64 `bson:"latency"`
	Ops     int64 `bson:"ops"`
}

// DBRecordStats stores data related to memory operations across databases.
type DBRecordStats struct {
	AccessesNotInMemory       int64                     `bson:"accessesNotInMemory"`
	PageFaultExceptionsThrown int64                     `bson:"pageFaultExceptionsThrown"`
	DBRecordAccesses          map[string]RecordAccesses `bson:",inline"`
}

// RecordAccesses stores data related to memory operations scoped to a database.
type RecordAccesses struct {
	AccessesNotInMemory       int64 `bson:"accessesNotInMemory"`
	PageFaultExceptionsThrown int64 `bson:"pageFaultExceptionsThrown"`
}

// MemStats stores data related to memory statistics.
type MemStats struct {
	Bits              int64       `bson:"bits"`
	Resident          int64       `bson:"resident"`
	Virtual           int64       `bson:"virtual"`
	Supported         interface{} `bson:"supported"`
	Mapped            int64       `bson:"mapped"`
	MappedWithJournal int64       `bson:"mappedWithJournal"`
}

// ReplStatus stores data related to replica sets.
type ReplStatus struct {
	SetName           string      `bson:"setName"`
	IsWritablePrimary interface{} `bson:"isWritablePrimary"` // mongodb 5.x
	IsMaster          interface{} `bson:"ismaster"`
	Secondary         interface{} `bson:"secondary"`
	IsReplicaSet      interface{} `bson:"isreplicaset"`
	ArbiterOnly       interface{} `bson:"arbiterOnly"`
	Hosts             []string    `bson:"hosts"`
	Passives          []string    `bson:"passives"`
	Me                string      `bson:"me"`
}

type StorageEngine struct {
	Name string `bson:"name"`
}

// WiredTiger stores information related to the WiredTiger storage engine.
type WiredTiger struct {
	Transaction TransactionStats       `bson:"transaction"`
	Concurrent  ConcurrentTransactions `bson:"concurrentTransactions"`
	Cache       CacheStats             `bson:"cache"`
	Connection  WTConnectionStats      `bson:"connection"`
	DataHandle  DataHandleStats        `bson:"data-handle"`
}

// TransactionStats stores transaction checkpoints in WiredTiger.
type TransactionStats struct {
	TransCheckpointsTotalTimeMsecs int64 `bson:"transaction checkpoint total time (msecs)"`
	TransCheckpoints               int64 `bson:"transaction checkpoints"`
}

type ConcurrentTransactions struct {
	Write ConcurrentTransStats `bson:"write"`
	Read  ConcurrentTransStats `bson:"read"`
}

type ConcurrentTransStats struct {
	Out          int64 `bson:"out"`
	Available    int64 `bson:"available"`
	TotalTickets int64 `bson:"totalTickets"`
}

// CacheStats stores cache statistics for WiredTiger.
type CacheStats struct {
	TrackedDirtyBytes         int64 `bson:"tracked dirty bytes in the cache"`
	CurrentCachedBytes        int64 `bson:"bytes currently in the cache"`
	MaxBytesConfigured        int64 `bson:"maximum bytes configured"`
	AppThreadsPageReadCount   int64 `bson:"application threads page read from disk to cache count"`
	AppThreadsPageReadTime    int64 `bson:"application threads page read from disk to cache time (usecs)"`
	AppThreadsPageWriteCount  int64 `bson:"application threads page write from cache to disk count"`
	AppThreadsPageWriteTime   int64 `bson:"application threads page write from cache to disk time (usecs)"`
	BytesWrittenFrom          int64 `bson:"bytes written from cache"`
	BytesReadInto             int64 `bson:"bytes read into cache"`
	PagesEvictedByAppThread   int64 `bson:"pages evicted by application threads"`
	PagesQueuedForEviction    int64 `bson:"pages queued for eviction"`
	PagesReadIntoCache        int64 `bson:"pages read into cache"`
	PagesWrittenFromCache     int64 `bson:"pages written from cache"`
	PagesRequestedFromCache   int64 `bson:"pages requested from the cache"`
	ServerEvictingPages       int64 `bson:"eviction server evicting pages"`
	WorkerThreadEvictingPages int64 `bson:"eviction worker thread evicting pages"`
	InternalPagesEvicted      int64 `bson:"internal pages evicted"`
	ModifiedPagesEvicted      int64 `bson:"modified pages evicted"`
	UnmodifiedPagesEvicted    int64 `bson:"unmodified pages evicted"`
}

// WTConnectionStats stores statistices on wiredTiger connections
type WTConnectionStats struct {
	FilesCurrentlyOpen int64 `bson:"files currently open"`
}

// DataHandleStats stores statistics for wiredTiger data-handles
type DataHandleStats struct {
	DataHandlesCurrentlyActive int64 `bson:"connection data handles currently active"`
}

// MetricsStats stores information related to metrics
type MetricsStats struct {
	TTL           *TTLStats           `bson:"ttl"`
	Cursor        *CursorStats        `bson:"cursor"`
	Document      *DocumentStats      `bson:"document"`
	Commands      *CommandsStats      `bson:"commands"`
	Operation     *OperationStats     `bson:"operation"`
	QueryExecutor *QueryExecutorStats `bson:"queryExecutor"`
	Repl          *ReplStats          `bson:"repl"`
	Storage       *StorageStats       `bson:"storage"`
}

// TTLStats stores information related to documents with a ttl index.
type TTLStats struct {
	DeletedDocuments int64 `bson:"deletedDocuments"`
	Passes           int64 `bson:"passes"`
}

// CursorStats stores information related to cursor metrics.
type CursorStats struct {
	TimedOut int64            `bson:"timedOut"`
	Open     *OpenCursorStats `bson:"open"`
}

// DocumentStats stores information related to document metrics.
type DocumentStats struct {
	Deleted  int64 `bson:"deleted"`
	Inserted int64 `bson:"inserted"`
	Returned int64 `bson:"returned"`
	Updated  int64 `bson:"updated"`
}

// CommandsStats stores information related to document metrics.
type CommandsStats struct {
	Aggregate     *CommandsStatsValue `bson:"aggregate"`
	Count         *CommandsStatsValue `bson:"count"`
	Delete        *CommandsStatsValue `bson:"delete"`
	Distinct      *CommandsStatsValue `bson:"distinct"`
	Find          *CommandsStatsValue `bson:"find"`
	FindAndModify *CommandsStatsValue `bson:"findAndModify"`
	GetMore       *CommandsStatsValue `bson:"getMore"`
	Insert        *CommandsStatsValue `bson:"insert"`
	Update        *CommandsStatsValue `bson:"update"`
}

type CommandsStatsValue struct {
	Failed int64 `bson:"failed"`
	Total  int64 `bson:"total"`
}

// OpenCursorStats stores information related to open cursor metrics
type OpenCursorStats struct {
	NoTimeout int64 `bson:"noTimeout"`
	Pinned    int64 `bson:"pinned"`
	Total     int64 `bson:"total"`
}

// OperationStats stores information related to query operations
// using special operation types
type OperationStats struct {
	ScanAndOrder   int64 `bson:"scanAndOrder"`
	WriteConflicts int64 `bson:"writeConflicts"`
}

// QueryExecutorStats stores information related to query execution
type QueryExecutorStats struct {
	Scanned        int64 `bson:"scanned"`
	ScannedObjects int64 `bson:"scannedObjects"`
}

// ReplStats stores information related to replication process
type ReplStats struct {
	Apply    *ReplApplyStats    `bson:"apply"`
	Buffer   *ReplBufferStats   `bson:"buffer"`
	Executor *ReplExecutorStats `bson:"executor,omitempty"`
	Network  *ReplNetworkStats  `bson:"network"`
}

// ReplApplyStats stores information related to oplog application process
type ReplApplyStats struct {
	Batches *BasicStats `bson:"batches"`
	Ops     int64       `bson:"ops"`
}

// ReplBufferStats stores information related to oplog buffer
type ReplBufferStats struct {
	Count     int64 `bson:"count"`
	SizeBytes int64 `bson:"sizeBytes"`
}

// ReplExecutorStats stores information related to replication executor
type ReplExecutorStats struct {
	Pool             map[string]int64 `bson:"pool"`
	Queues           map[string]int64 `bson:"queues"`
	UnsignaledEvents int64            `bson:"unsignaledEvents"`
}

// ReplNetworkStats stores information related to network usage by replication process
type ReplNetworkStats struct {
	Bytes    int64       `bson:"bytes"`
	GetMores *BasicStats `bson:"getmores"`
	Ops      int64       `bson:"ops"`
}

// BasicStats stores information about an operation
type BasicStats struct {
	Num         int64 `bson:"num"`
	TotalMillis int64 `bson:"totalMillis"`
}

// StorageStats stores information related to record allocations
type StorageStats struct {
	FreelistSearchBucketExhausted int64 `bson:"freelist.search.bucketExhausted"`
	FreelistSearchRequests        int64 `bson:"freelist.search.requests"`
	FreelistSearchScanned         int64 `bson:"freelist.search.scanned"`
}

// TCMallocStats stores information related to TCMalloc memory allocator metrics
type TCMallocStats struct {
	Generic  *GenericTCMAllocStats  `bson:"generic"`
	TCMalloc *DetailedTCMallocStats `bson:"tcmalloc"`
}

// GenericTCMAllocStats stores generic TCMalloc memory allocator metrics
type GenericTCMAllocStats struct {
	CurrentAllocatedBytes int64 `bson:"current_allocated_bytes"`
	HeapSize              int64 `bson:"heap_size"`
}

// DetailedTCMallocStats stores detailed TCMalloc memory allocator metrics
type DetailedTCMallocStats struct {
	PageheapFreeBytes            int64 `bson:"pageheap_free_bytes"`
	PageheapUnmappedBytes        int64 `bson:"pageheap_unmapped_bytes"`
	MaxTotalThreadCacheBytes     int64 `bson:"max_total_thread_cache_bytes"`
	CurrentTotalThreadCacheBytes int64 `bson:"current_total_thread_cache_bytes"`
	TotalFreeBytes               int64 `bson:"total_free_bytes"`
	CentralCacheFreeBytes        int64 `bson:"central_cache_free_bytes"`
	TransferCacheFreeBytes       int64 `bson:"transfer_cache_free_bytes"`
	ThreadCacheFreeBytes         int64 `bson:"thread_cache_free_bytes"`
	PageheapComittedBytes        int64 `bson:"pageheap_committed_bytes"`
	PageheapScavengeCount        int64 `bson:"pageheap_scavenge_count"`
	PageheapCommitCount          int64 `bson:"pageheap_commit_count"`
	PageheapTotalCommitBytes     int64 `bson:"pageheap_total_commit_bytes"`
	PageheapDecommitCount        int64 `bson:"pageheap_decommit_count"`
	PageheapTotalDecommitBytes   int64 `bson:"pageheap_total_decommit_bytes"`
	PageheapReserveCount         int64 `bson:"pageheap_reserve_count"`
	PageheapTotalReserveBytes    int64 `bson:"pageheap_total_reserve_bytes"`
	SpinLockTotalDelayNanos      int64 `bson:"spinlock_total_delay_ns"`
}

type CollectionStats struct {
	Collection     string  `bson:"ns"`
	Count          int64   `bson:"count"`
	Size           int64   `bson:"size"`
	AvgObjSize     float64 `bson:"avgObjSize"`
	StorageSize    int64   `bson:"storageSize"`
	TotalIndexSize int64   `bson:"totalIndexSize"`
	Ok             int64   `bson:"ok"`
}
