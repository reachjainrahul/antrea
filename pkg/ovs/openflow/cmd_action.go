package openflow

import (
	"fmt"
	"strings"
)

type commandAction struct {
	builder *commandBuilder
}

func (a *commandAction) Drop() FlowBuilder {
	a.builder.actions = append(a.builder.actions, "drop")
	return a.builder
}

func (a *commandAction) Output(port int) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("output:%d", port))
	return a.builder
}

func (a *commandAction) OutputFieldRange(name string, rng Range) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("output:%s[%d..%d]", name, rng[0], rng[1]))
	return a.builder
}

func (a *commandAction) OutputInPort() FlowBuilder {
	a.builder.actions = append(a.builder.actions, "in_port")
	return a.builder
}

func (a *commandAction) CT(commit bool, tableID TableIDType, zone int, actions ...string) FlowBuilder {
	var repr string
	if commit {
		repr += "commit"
	}
	if tableID > 0 {
		if repr != "" {
			repr += ","
		}
		repr += fmt.Sprintf("table=%d", tableID)
	}
	if zone > 0 {
		if repr != "" {
			repr += ","
		}
		repr += fmt.Sprintf("zone=%d", zone)
	}
	if len(actions) > 0 {
		if repr != "" {
			repr += ","
		}
		repr += fmt.Sprintf("exec(%s)", strings.Join(actions, ","))
	}
	a.builder.actions = append(a.builder.actions, "ct("+repr+")")
	return a.builder
}

func (a *commandAction) SetField(key, value string) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("set_field:%s->%s", value, key))
	return a.builder
}

func (a *commandAction) Load(name string, value uint64) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("load:0x%x->%s[]", value, name))
	return a.builder
}

func (a *commandAction) LoadRange(name string, addr uint32, to Range) FlowBuilder {
	a.builder.actions = append(
		a.builder.actions,
		fmt.Sprintf("load:0x%x->%s[%d..%d]", addr, name, to[0], to[1]),
	)
	return a.builder
}

func (a *commandAction) Move(from, to string) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("move:%s[]->%s[]", from, to))
	return a.builder
}

func (a *commandAction) MoveRange(fromName, toName string, from, to Range) FlowBuilder {
	a.builder.actions = append(
		a.builder.actions,
		fmt.Sprintf("move:%s[%d..%d]->%s[%d..%d]", fromName, from[0], from[1], toName, to[0], to[1]),
	)
	return a.builder
}

func (a *commandAction) Resubmit(port string, table TableIDType) FlowBuilder {
	a.builder.actions = append(a.builder.actions, fmt.Sprintf("resubmit(%s,%d)", port, table))
	return a.builder
}

func (a *commandAction) DecTTL() FlowBuilder {
	a.builder.actions = append(a.builder.actions, "dec_ttl")
	return a.builder
}

func (a *commandAction) Normal() FlowBuilder {
	a.builder.actions = append(a.builder.actions, "Normal")
	return a.builder
}
