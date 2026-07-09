# GO_RUNBOOK

## Understanding Go for DevOps & Platform Engineering

**Author:** Raunak Pandey
**Project Context:** NorthStar – Engineering Intelligence Platform

---

# 1. Why Go?

NorthStar is an Engineering Intelligence Platform that continuously collects
data from multiple external systems including GitHub, Jenkins, Jira,
Kubernetes, and Prometheus.

These integrations are largely independent and network-bound.

Go's lightweight goroutines and GMP scheduler allow the platform to perform
thousands of concurrent I/O operations while maintaining low memory usage and
high throughput, making it well suited for cloud-native backend services.

Instead of focusing only on writing applications, Go was designed for:

* Cloud Computing
* Networking
* Distributed Systems
* Infrastructure Tools
* Platform Engineering
* DevOps

This is why many cloud-native tools are written in Go.

Examples:

* Kubernetes
* Docker
* Prometheus
* Terraform
* Helm
* Caddy
* Grafana Agent

Go is not "better" than Python.

It simply solves different engineering problems.

---

# 2. Why NorthStar uses Go

NorthStar constantly communicates with multiple systems.

Example:

GitHub

↓

Jenkins

↓

GitLab

↓

Jira

↓

Kubernetes

↓

Prometheus

↓

PostgreSQL

Every API call is independent.

Instead of waiting for one API before calling another, Go allows them to execute concurrently.

This significantly improves throughput.

---

# 3. Process vs Thread vs Goroutine

## Process

A running program.

Examples:

* Chrome
* VS Code
* NorthStar

Each process has its own memory.

---

## Thread

A worker inside a process.

Threads are created and managed by the Operating System.

Characteristics:

* Heavyweight
* Large memory usage
* Expensive to create
* Expensive to switch

---

## Goroutine

A lightweight task managed by the Go Runtime.

Characteristics:

* Extremely lightweight
* Starts with a very small stack
* Thousands can exist simultaneously
* Scheduled by Go instead of the Operating System

A goroutine is **NOT** an OS thread.

---

# 4. Concurrency vs Parallelism

## Concurrency

Making progress on multiple tasks.

One CPU core can switch between tasks.

Example:

Task A

↓

Task B

↓

Task C

↓

Back to Task A

---

## Parallelism

Multiple tasks actually execute simultaneously.

Requires multiple CPU cores.

Example:

Core 1 → Task A

Core 2 → Task B

Core 3 → Task C

---

Go provides a simple concurrency model using goroutines.

If multiple CPU cores exist, Go can also execute goroutines in parallel.

---

# 5. The Go Runtime

The Go Runtime is an internal system that manages the execution of Go programs.

It provides:

* GMP Scheduler
* Garbage Collector
* Memory Allocator
* Stack Management
* Network Poller
* Timers
* Panic/Recover Support

Think of it as a miniature operating system inside every Go application.

---

# 6. GMP Scheduler

G = Goroutine

M = Machine

P = Processor

---

## G — Goroutine

A lightweight task.

Example:

go FetchGitHub()

go FetchJenkins()

go CalculateDORA()

---

## M — Machine

An Operating System Thread.

Created and scheduled by the OS.

---

## P — Processor

A logical scheduling context managed by Go.

Important:

P is **NOT** a physical CPU.

It owns a local run queue of runnable goroutines.

A Machine (OS Thread) must acquire a Processor before executing goroutines.

---

# 7. How Scheduling Works

Your Code

↓

Creates Goroutines

↓

Go Runtime Scheduler (GMP)

↓

OS Threads

↓

CPU

The Operating System schedules threads.

The Go Runtime schedules goroutines.

These are two separate scheduling layers.

---

# 8. Waiting Goroutines

Example:

Fetch GitHub

↓

GitHub API

↓

Waiting...

The goroutine becomes blocked.

The OS thread does **NOT** remain idle.

Instead, the Go Runtime immediately schedules another runnable goroutine on that thread.

When GitHub responds:

Waiting

↓

Runnable

↓

Placed back into a run queue

↓

Executed by whichever thread becomes available

A goroutine is **not permanently attached** to any specific OS thread.

---

# 9. Work Stealing

Each Processor has its own local queue.

If one Processor has no work:

Processor B

↓

Steals runnable goroutines

↓

From Processor A

This keeps CPU utilization high.

---

# 10. Why Goroutines are Powerful

Instead of:

100,000 OS Threads

Go uses:

100,000 Goroutines

↓

Small Number of OS Threads

↓

CPU

Benefits:

* Low memory usage
* Lower scheduling overhead
* Excellent scalability
* Efficient handling of I/O-bound workloads

---

# 11. Does Every Goroutine Run Simultaneously?

No.

Example:

Laptop

6 Cores

12 Logical CPUs

Only about 12 goroutines can execute CPU instructions simultaneously.

However,

100,000+ goroutines can:

* Exist
* Wait
* Sleep
* Become Runnable
* Be Scheduled

The Go Runtime rapidly switches between them.

---

# 12. Why Go is Used for DevOps

Go excels at:

* API Servers
* Monitoring Agents
* Kubernetes Controllers
* Docker Internals
* Infrastructure Tools
* CI/CD Systems
* Network Services
* Cloud-Native Platforms

Most DevOps workloads spend significant time waiting for:

* APIs
* Databases
* Files
* Networks

Go keeps CPU resources productive during these waits.

---

# 13. Why "Docker is Written in Go" Matters

This does **NOT** mean Docker commands become easier.

It means:

* Docker internals are written in Go.
* Docker extensions and related libraries are Go-based.
* Reading Docker source becomes easier.
* Contributing to Docker is easier.
* Understanding Kubernetes, Terraform, and Prometheus internals becomes more natural.

---

# 14. Compiled vs Interpreted

Python:

Source Code

↓

Interpreter

↓

Machine Instructions

↓

CPU

Go:

Source Code

↓

Compiler

↓

Native Binary

↓

CPU

Go executes native machine code directly.

This reduces runtime overhead.

However, for network-heavy applications, the largest delay usually comes from waiting on external systems rather than language execution speed.

---

# 15. NorthStar Architecture Decision

Go Responsibilities

* REST APIs
* GitHub Integration
* Jenkins Integration
* GitLab Integration
* Kubernetes Integration
* Collector Engine
* Scheduler
* Analytics Engine
* DORA Metric Calculation
* Background Workers

Python Responsibilities (Future)

* AI Summaries
* Predictive Analytics
* Trend Analysis
* LLM Integration

---

# 16. Interview Nuggets

**Why Go?**

"I chose Go because NorthStar performs many independent API calls to GitHub, Jenkins, Jira, and Kubernetes. Goroutines and the GMP scheduler allow these I/O-bound operations to run concurrently with low overhead, making the collector service scalable and resource-efficient."

---

**Why not one thread per request?**

"OS threads are heavyweight. Go multiplexes many lightweight goroutines onto a smaller number of OS threads, reducing memory usage and scheduling overhead."

---

**Can a goroutine resume on another thread?**

Yes.

A blocked goroutine is not tied to any specific OS thread. Once it becomes runnable again, the Go scheduler may resume it on whichever thread is available.

---

# 17. Mental Models to Remember

## Company Analogy

Company = Process

Employee = OS Thread

Task = Goroutine

Desk = Processor (P)

Manager = Go Runtime Scheduler

---

## Courier Analogy

Packages = Goroutines

Delivery Trucks = OS Threads

Dispatcher = Go Runtime

Roads = CPU Cores

The dispatcher constantly assigns available trucks to pending deliveries.

---

# 18. Key Takeaways

* Goroutines are lightweight tasks, not threads.
* The Go Runtime schedules goroutines; the OS schedules threads.
* Goroutines are not permanently attached to threads.
* A blocked goroutine can resume on any available thread.
* GMP is the scheduler subsystem inside the Go Runtime.
* Go excels at concurrent, I/O-heavy systems.
* NorthStar is an ideal Go project because it aggregates data from many independent sources concurrently.

---

## Final Thought

> **Go isn't powerful because it's just fast. It's powerful because it lets developers think in terms of independent tasks while the runtime efficiently manages scheduling, concurrency, and resource utilization.**
