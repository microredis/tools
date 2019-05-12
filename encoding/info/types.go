package info

type DB struct {
	Expires int64 `json:"expires,string"`
	AvgTTL  int64 `json:"avg_ttl,string"`
	Keys    int64 `json:"keys,string"`
}

type Cmdstat struct {
	UsecPerCall float32 `json:"usec_per_call,string"`
	Usec        int     `json:"usec,string"`
	Calls       int     `json:"calls,string"`
}

type Stats struct {
	RedisVersion             string  `json:"redis_version"`
	ConnectedClients         int64   `json:"connected_clients,string,omitempty"`
	UsedCpuSys               float64 `json:"used_cpu_sys,string,omitempty"`
	UsedCpuUser              float64 `json:"used_cpu_user,string,omitempty"`
	UsedMemory               int64   `json:"used_memory,string,omitempty"`
	Maxmemory                int64   `json:"maxmemory,string,omitempty"`
	TotalConnectionsReceived int64   `json:"total_connections_received,string,omitempty"`
	TotalCommandsProcessed   int64   `json:"total_commands_processed,string,omitempty"`
	ExpiredKeys              int64   `json:"expired_keys,string,omitempty"`
	EvictedKeys              int64   `json:"evicted_keys,string,omitempty"`
	KeyspaceHits             int64   `json:"keyspace_hits,string,omitempty"`
	KeyspaceMisses           int64   `json:"keyspace_misses,string,omitempty"`
	PubsubChannels           int64   `json:"pubsub_channels,string,omitempty"`
	PubsubPatterns           int64   `json:"pubsub_patterns,string,omitempty"`
	UsedMemoryLua            int64   `json:"used_memory_lua,string,omitempty"`
	UsedMemoryPeak           int64   `json:"used_memory_peak,string,omitempty"`
	UsedMemoryRSS            int64   `json:"used_memory_rss,string,omitempty"`
}

type Info struct {
	Stats    `json:"stats"`
	Keyspace map[string]*DB      `json:"keyspace"`
	Cmdstats map[string]*Cmdstat `json:"cmdstats"`
}
