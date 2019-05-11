package info

import (
	"strings"
	"testing"
)

type MyStruct struct {
	RedisVersion string                      `info:"redis_version"`
	Keyspace     map[string]map[string]int64 `info:",keyspace"`
	CmdstatLlen  map[string]string           `info:"cmdstat_llen"`
}

func TestUnmarshal(t *testing.T) {
	data := strings.Join(strings.Split(data, "\n"), "\r\n")
	myStruct := new(MyStruct)
	if err := Unmarshal([]byte(data), myStruct); err != nil {
		t.Fatal("failed", err.Error())
	}
	if myStruct.RedisVersion != "999.999.999" {
		t.Fatal("redis version mismatch with ", myStruct.RedisVersion)
	}
	if myStruct.Keyspace["db0"] == nil {
		t.Fatal("keyspace is not parsed")
	}
	if myStruct.Keyspace["db0"]["keys"] != 2528768 {
		t.Fatal("keyspace is not parsed")
	}
	if myStruct.Keyspace["db0"]["expires"] != 1420 {
		t.Fatal("keyspace is not parsed")
	}
	if myStruct.Keyspace["db0"]["avg_ttl"] != 107761233218 {
		t.Fatal("keyspace is not parsed")
	}
	if myStruct.CmdstatLlen["calls"] != "5" {
		t.Fatal("CmdstatLlen is not parsed")
	}
	if myStruct.CmdstatLlen["usec"] != "8" {
		t.Fatal("CmdstatLlen is not parsed")
	}
	if myStruct.CmdstatLlen["usec_per_call"] != "1.60" {
		t.Fatal("CmdstatLlen is not parsed")
	}
}

const data = `
# Server
redis_version:999.999.999
redis_git_sha1:3c968ff0
redis_git_dirty:0
redis_build_id:51089de051945df4
redis_mode:standalone
os:Linux 4.8.0-1-amd64 x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:atomic-builtin
gcc_version:6.3.0
process_id:3036
run_id:868be887c3c27b100329d99321454df44d1e5394
tcp_port:6379
uptime_in_seconds:4863493
uptime_in_days:56
hz:10
lru_clock:3626017
executable:/usr/local/bin/redis-server
config_file:

# Clients
connected_clients:4
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:0

# Memory
used_memory:441315472
used_memory_human:420.87M
used_memory_rss:454508544
used_memory_rss_human:433.45M
used_memory_peak:441315472
used_memory_peak_human:420.87M
used_memory_peak_perc:100.00%
used_memory_overhead:135465032
used_memory_startup:510704
used_memory_dataset:305850440
used_memory_dataset_perc:69.38%
allocator_allocated:441349720
allocator_active:441700352
allocator_resident:453648384
total_system_memory:1044770816
total_system_memory_human:996.37M
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:350632
allocator_rss_ratio:1.03
allocator_rss_bytes:11948032
rss_overhead_ratio:1.00
rss_overhead_bytes:860160
mem_fragmentation_ratio:1.03
mem_fragmentation_bytes:13235008
mem_allocator:jemalloc-4.0.3
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:12555661
rdb_bgsave_in_progress:0
rdb_last_save_time:1542266396
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:0
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:0

# Stats
total_connections_received:70
total_commands_processed:31140633
instantaneous_ops_per_sec:16
total_net_input_bytes:2572538317
total_net_output_bytes:435242454
instantaneous_input_kbps:1.32
instantaneous_output_kbps:0.17
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:39697
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:7257811
keyspace_misses:3078844
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:0
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:0
master_replid:158d8d877375ee481819a87b83da988bc92e5b23
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:3349.96
used_cpu_user:37121.75
used_cpu_sys_children:0.00
used_cpu_user_children:0.00

# Cluster
cluster_enabled:0

# Commandstats
cmdstat_llen:calls=5,usec=8,usec_per_call=1.60
cmdstat_slowlog:calls=923,usec=2097,usec_per_call=2.27
cmdstat_set:calls=2,usec=9,usec_per_call=4.50
cmdstat_strlen:calls=25,usec=49,usec_per_call=1.96
cmdstat_command:calls=1,usec=4474,usec_per_call=4474.00
cmdstat_type:calls=273,usec=759,usec_per_call=2.78
cmdstat_auth:calls=275,usec=734,usec_per_call=2.67
cmdstat_info:calls=1848,usec=72492,usec_per_call=39.23
cmdstat_select:calls=6,usec=17,usec_per_call=2.83
cmdstat_xinfo:calls=7,usec=95,usec_per_call=13.57
cmdstat_lrange:calls=3,usec=633,usec_per_call=211.00
cmdstat_config:calls=923,usec=52759,usec_per_call=57.16
cmdstat_xlen:calls=239,usec=465,usec_per_call=1.95
cmdstat_scan:calls=9,usec=63,usec_per_call=7.00
cmdstat_lpush:calls=1,usec=939,usec_per_call=939.00
cmdstat_get:calls=1,usec=9,usec_per_call=9.00
cmdstat_client:calls=262,usec=355,usec_per_call=1.35
cmdstat_xadd:calls=28,usec=3957,usec_per_call=141.32
cmdstat_ping:calls=927,usec=548,usec_per_call=0.59

# Keyspace
db0:keys=2528768,expires=1420,avg_ttl=107761233218`
