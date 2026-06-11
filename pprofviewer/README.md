# pprofviewer

Independent web service for quickly inspecting Go heap and CPU pprof files.

## Run

```bash
cd pprofviewer
GOWORK=off go run ./cmd/pprofviewer -addr :7777
```

Open `http://localhost:7777`, upload a heap or CPU profile, and choose the sample type. Heap profiles default to `alloc_space`; CPU profiles default to `samples` or `cpu` when available.

From the repository root:

```bash
make pprofviewer
make pprofviewer PPROF_FILES="/tmp/heap.pprof /tmp/cpu.pprof"
make pprofviewer PPROF_HEAP=/tmp/heap.pprof PPROF_CPU=/tmp/cpu.pprof
```

The Makefile recipe prints direct viewer URLs for each supplied profile path. Direct URLs also render automatically:

```text
http://localhost:7777/?path=/tmp/heap.pprof&sample=inuse_space
```

## API

Analyze a local file from the server process:

```bash
curl "http://localhost:7777/api/analyze?path=/tmp/heap.pprof&sample=inuse_space"
```

Analyze an uploaded profile:

```bash
curl -F "profile=@/tmp/cpu.pprof" "http://localhost:7777/api/analyze"
```

The response includes:

- `nodes`: graph nodes sized by cumulative allocation or CPU sample value
- `edges`: weighted call-path edges
- `leaks`: retained heap candidates with stack paths
- `top_paths`: heaviest stack paths in the profile
