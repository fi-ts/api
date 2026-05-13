import type { GenExtension, GenFile } from "@bufbuild/protobuf/codegenv2";
import type { MapRules, RepeatedRules, StringRules } from "../../../buf/validate/validate_pb";
/**
 * Describes the file fits/api/v1/predefined_rules.proto.
 */
export declare const file_fits_api_v1_predefined_rules: GenFile;
/**
 * Macaddress returns true if the given string is a valid macadress
 *
 * @generated from extension: optional bool macaddress = 80048951;
 */
export declare const macaddress: GenExtension<StringRules, boolean>;
/**
 * IsName returns true if name field satisfies our requirements
 *
 * @generated from extension: optional bool is_name = 80048952;
 */
export declare const is_name: GenExtension<StringRules, boolean>;
/**
 * IsDescription returns true if description field satisfies our requirements
 *
 * @generated from extension: optional bool is_description = 80048953;
 */
export declare const is_description: GenExtension<StringRules, boolean>;
/**
 * IsPartition returns true if partition field satisfies our requirements
 *
 * @generated from extension: optional bool is_partition = 80048954;
 */
export declare const is_partition: GenExtension<StringRules, boolean>;
/**
 * IsPrefix validates if the given string is a valid prefix
 *
 * @generated from extension: optional bool is_prefix = 80048955;
 */
export declare const is_prefix: GenExtension<StringRules, boolean>;
/**
 * IsUri validates if the given string is a valid URI
 *
 * @generated from extension: optional bool is_uri = 80048956;
 */
export declare const is_uri: GenExtension<StringRules, boolean>;
/**
 * IsIpOrHostname validates that the given string is either a ip or a hostname
 *
 * @generated from extension: optional bool is_ip_or_hostname = 80048957;
 */
export declare const is_ip_or_hostname: GenExtension<StringRules, boolean>;
/**
 * Trimmed enforces the string to be trimmed, e.g. no whitespaces at the begin and end
 *
 * @generated from extension: optional bool trimmed = 80048958;
 */
export declare const trimmed: GenExtension<StringRules, boolean>;
/**
 * Prefixes validates if a slice of prefixes in string form are valid
 *
 * @generated from extension: optional bool prefixes = 80058951;
 */
export declare const prefixes: GenExtension<RepeatedRules, boolean>;
/**
 * Ips validates if a slice of ips in string form are valid
 *
 * @generated from extension: optional bool ips = 80058952;
 */
export declare const ips: GenExtension<RepeatedRules, boolean>;
/**
 * AreHostAndPort validates if a slice of strings are all in the form of <ip | host>:<port>
 *
 * @generated from extension: optional bool are_host_and_port = 80058953;
 */
export declare const are_host_and_port: GenExtension<RepeatedRules, boolean>;
/**
 * All Trimmed enforces all strings to be trimmed, e.g. no whitespaces at the begin and end
 *
 * @generated from extension: optional bool all_trimmed = 80058954;
 */
export declare const all_trimmed: GenExtension<RepeatedRules, boolean>;
/**
 * Keys and Values trimmed enforces all map keys and values to be trimmed, e.g. no whitespaces at the begin and end
 *
 * @generated from extension: optional bool keys_and_values_trimmed = 80068951;
 */
export declare const keys_and_values_trimmed: GenExtension<MapRules, boolean>;
/**
 * Keys trimmed enforces all map keys and values to be trimmed, e.g. no whitespaces at the begin and end
 *
 * @generated from extension: optional bool keys_trimmed = 80068952;
 */
export declare const keys_trimmed: GenExtension<MapRules, boolean>;
