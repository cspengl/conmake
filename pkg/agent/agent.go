package agent

type Agent interface {
  PerformStep() void
  Info() void
}
