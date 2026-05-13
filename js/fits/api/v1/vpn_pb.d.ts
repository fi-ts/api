import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";
/**
 * Describes the file fits/api/v1/vpn.proto.
 */
export declare const file_fits_api_v1_vpn: GenFile;
/**
 * VPNNode represents a machine connected to the VPN.
 *
 * @generated from message fits.api.v1.VPNNode
 */
export type VPNNode = Message<"fits.api.v1.VPNNode"> & {
    /**
     * ID of this node
     *
     * @generated from field: uint64 id = 1;
     */
    id: bigint;
    /**
     * Name of this node
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * Project of this node, maps to a project
     *
     * @generated from field: string project = 3;
     */
    project: string;
    /**
     * IPAddresses of this node in the VPN
     *
     * @generated from field: repeated string ip_addresses = 4;
     */
    ipAddresses: string[];
    /**
     * LastSeen timestamp when this node reached out to the control plane
     *
     * @generated from field: google.protobuf.Timestamp last_seen = 5;
     */
    lastSeen?: Timestamp | undefined;
    /**
     * Online indicates if this node is online
     *
     * @generated from field: bool online = 6;
     */
    online: boolean;
};
/**
 * Describes the message fits.api.v1.VPNNode.
 * Use `create(VPNNodeSchema)` to create a new message.
 */
export declare const VPNNodeSchema: GenMessage<VPNNode>;
