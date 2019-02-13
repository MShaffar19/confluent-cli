package kafka

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	kafkav1 "github.com/confluentinc/ccloudapis/kafka/v1"
)

// ACLConfiguration wrapper used for flag parsing and validation
type ACLConfiguration struct {
	*kafkav1.ACLBinding
	errors []string
}

// aclConfigFlags returns a flag set which can be parsed to create an ACLConfiguration object.
func aclConfigFlags() *pflag.FlagSet {
	flgSet := aclEntryFlags()
	flgSet.SortFlags = false
	flgSet.AddFlagSet(resourceFlags())
	return flgSet
}

// aclEntryFlags returns a flag set which can be parsed to create an AccessControlEntry object.
func aclEntryFlags() *pflag.FlagSet {
	flgSet := pflag.NewFlagSet("acl-entry", pflag.ExitOnError)
	flgSet.Bool("allow", false, "Set ACL to grant access")
	flgSet.Bool("deny", false, "Set ACL to restrict access to resource")
	//flgSet.String( "host", "*", "Set Kafka principal host. Note: Not supported on CCLOUD.")
	flgSet.String("principal", "", "Set Kafka principal")
	flgSet.String("operation", "", fmt.Sprintf("Set ACL Operation: [%s]",
		listEnum(kafkav1.ACLOperations_ACLOperation_name, []string{"ANY", "UNKNOWN"})))
	// An error is only returned if the flag name is not present.
	// We know the flag name is present so its safe to ignore this.
	_ = cobra.MarkFlagRequired(flgSet, "principal")
	_ = cobra.MarkFlagRequired(flgSet, "operation")
	return flgSet
}

// resourceFlags returns a flag set which can be parsed to create a ResourcePattern object.
func resourceFlags() *pflag.FlagSet {
	flgSet := pflag.NewFlagSet("acl-resource", pflag.ExitOnError)
	flgSet.Bool("cluster", false, "Set CLUSTER resource")
	flgSet.String("topic", "", "Set TOPIC resource")
	flgSet.String("consumer-group", "", "Set CONSUMER_GROUP resource")
	flgSet.String("transactional-id", "", "Set TRANSACTIONAL_ID resource")
	flgSet.Bool("prefix", false, "Set to match all resource names prefixed with this value")

	return flgSet
}

// parse returns ACLConfiguration from the contents of cmd
func parse(cmd *cobra.Command) *ACLConfiguration {
	aclBinding := &ACLConfiguration{
		ACLBinding: &kafkav1.ACLBinding{
			Entry: &kafkav1.AccessControlEntryConfig{
				Host: "*",
			},
		},
	}
	cmd.Flags().Visit(fromArgs(aclBinding))
	return aclBinding
}

// fromArgs maps command flag values to the appropriate ACLConfiguration field
func fromArgs(conf *ACLConfiguration) func(*pflag.Flag) {
	return func(flag *pflag.Flag) {
		v := flag.Value.String()
		switch n := flag.Name; n {
		case "consumer-group":
			setResourcePattern(conf, "GROUP", v)
		case "cluster":
			// The only valid name for a cluster is kafka-cluster
			// https://github.com/confluentinc/cc-kafka/blob/88823c6016ea2e306340938994d9e122abf3c6c0/core/src/main/scala/kafka/security/auth/Resource.scala#L24
			setResourcePattern(conf, n, "kafka-cluster")
		case "topic":
			fallthrough
		case "delegation-token":
			fallthrough
		case "transactional-id":
			setResourcePattern(conf, n, v)
		case "allow":
			conf.Entry.PermissionType = kafkav1.ACLPermissionTypes_ALLOW
		case "deny":
			conf.Entry.PermissionType = kafkav1.ACLPermissionTypes_DENY
		case "prefix":
			conf.Pattern.PatternType = kafkav1.PatternTypes_PREFIXED
		case "principal":
			conf.Entry.Principal = "User:" + v
		case "operation":
			v = strings.ToUpper(v)
			if op, ok := kafkav1.ACLOperations_ACLOperation_value[v]; ok {
				conf.Entry.Operation = kafkav1.ACLOperations_ACLOperation(op)
				break
			}
			conf.errors = append(conf.errors, "Invalid operation value: "+v)
		}
	}
}

func setResourcePattern(conf *ACLConfiguration, n, v string) {
	/* Normalize the resource pattern name */
	n = strings.ToUpper(n)
	n = strings.Replace(n, "-", "_", -1)

	if conf.Pattern != nil {
		conf.errors = append(conf.errors, "only one resource can be specified per command execution")
		return
	}

	conf.Pattern = &kafkav1.ResourcePatternConfig{}
	conf.Pattern.ResourceType = kafkav1.ResourceTypes_ResourceType(kafkav1.ResourceTypes_ResourceType_value[n])

	conf.Pattern.Name = v
	conf.Pattern.PatternType = kafkav1.PatternTypes_LITERAL
}

func listEnum(enum map[int32]string, exclude []string) string {
	var ops []string

OUTER:
	for _, v := range enum {
		for _, exclusion := range exclude {
			if v == exclusion {
				continue OUTER
			}
		}
		ops = append(ops, v)
	}

	return strings.Join(ops, ", ")
}
