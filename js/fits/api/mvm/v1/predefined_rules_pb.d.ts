import type { GenExtension, GenFile } from "@bufbuild/protobuf/codegenv2";
import type { RepeatedRules, StringRules, UInt32Rules, UInt64Rules } from "../../../../buf/validate/validate_pb.js";
/**
 * Describes the file fits/api/mvm/v1/predefined_rules.proto.
 */
export declare const file_fits_api_mvm_v1_predefined_rules: GenFile;
/**
 * IsNameCode validates if the given string is a 4-character code consisting of digits and letters.
 *
 * @generated from extension: optional bool is_name_code = 80108951;
 */
export declare const is_name_code: GenExtension<StringRules, boolean>;
/**
 * IsFqdn validates if the given string is a valid fully qualified domain name (FQDN),
 * e.g. host.example.com. Each label is alphanumeric/hyphen, the TLD is alphabetic
 * (min 2 chars), and the total length is at most 253 characters.
 *
 * @generated from extension: optional bool is_fqdn = 80108952;
 */
export declare const is_fqdn: GenExtension<StringRules, boolean>;
/**
 * IsOrderNumber validates a user-supplied ticket/reference ID (order number):
 * 2-128 characters, no leading/trailing whitespace.
 *
 * @generated from extension: optional bool is_order_number = 80108953;
 */
export declare const is_order_number: GenExtension<StringRules, boolean>;
/**
 * IsMountPoint validates a Linux mount point: an absolute path without
 * whitespace, e.g. /data/app. Allowed paths are enforced upstream.
 *
 * @generated from extension: optional bool is_mount_point = 80108954;
 */
export declare const is_mount_point: GenExtension<StringRules, boolean>;
/**
 * IsDriveLetter validates a Windows drive letter: a single letter, e.g. G.
 *
 * @generated from extension: optional bool is_drive_letter = 80108955;
 */
export declare const is_drive_letter: GenExtension<StringRules, boolean>;
/**
 * IsSnapshotName validates a snapshot name: 1-128 alphanumeric characters.
 *
 * @generated from extension: optional bool is_snapshot_name = 80108956;
 */
export declare const is_snapshot_name: GenExtension<StringRules, boolean>;
/**
 * IsCpuCount validates the number of CPU cores of a MVM.
 *
 * @generated from extension: optional bool is_cpu_count = 80078951;
 */
export declare const is_cpu_count: GenExtension<UInt32Rules, boolean>;
/**
 * IsRamSize validates the RAM size of a MVM in GiB.
 *
 * @generated from extension: optional bool is_ram_size = 80078952;
 */
export declare const is_ram_size: GenExtension<UInt32Rules, boolean>;
/**
 * IsDiskSize validates the size of a single MVM disk in GiB.
 *
 * @generated from extension: optional bool is_disk_size = 80088951;
 */
export declare const is_disk_size: GenExtension<UInt64Rules, boolean>;
/**
 * IsDiskCount validates the number of disks in a MVM request (at most 8).
 *
 * @generated from extension: optional bool is_disk_count = 80098951;
 */
export declare const is_disk_count: GenExtension<RepeatedRules, boolean>;
