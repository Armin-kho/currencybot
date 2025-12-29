package bot

import "github.com/shirou/gopsutil/process"

func (b *Bot) getStats() string {
	p, _ := process.NewProcess(int32(os.Getpid()))
	mem, _ := p.MemoryInfo()
	cpu, _ := p.CPUPercent()
	return fmt.Sprintf("CPU%%: %.2f\nMemory: %v\n",
		cpu, mem.RSS)
}
