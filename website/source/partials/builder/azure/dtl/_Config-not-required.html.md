<!-- Code generated from the comments of the Config struct in builder/azure/dtl/config.go; DO NOT EDIT MANUALLY -->

-   `capture_name_prefix` (string) - Capture
    
-   `capture_container_name` (string) - Capture Container Name
-   `shared_image_gallery` (SharedImageGallery) - Use a [Shared Gallery
    image](https://azure.microsoft.com/en-us/blog/announcing-the-public-preview-of-shared-image-gallery/)
    as the source for this build. *VHD targets are incompatible with this
    build type* - the target must be a *Managed Image*.
    
        "shared_image_gallery": {
            "subscription": "00000000-0000-0000-0000-00000000000",
            "resource_group": "ResourceGroup",
            "gallery_name": "GalleryName",
            "image_name": "ImageName",
            "image_version": "1.0.0"
        }
        "managed_image_name": "TargetImageName",
        "managed_image_resource_group_name": "TargetResourceGroup"
    
    
-   `shared_image_gallery_destination` (SharedImageGalleryDestination) - The name of the Shared Image Gallery under which the managed image will be published as Shared Gallery Image version.
    
    Following is an example.
    
    <!-- -->
    
        "shared_image_gallery_destination": {
            "resource_group": "ResourceGroup",
            "gallery_name": "GalleryName",
            "image_name": "ImageName",
            "image_version": "1.0.0",
            "replication_regions": ["regionA", "regionB", "regionC"]
        }
        "managed_image_name": "TargetImageName",
        "managed_image_resource_group_name": "TargetResourceGroup"
    
    
-   `shared_image_gallery_timeout` (time.Duration) - How long to wait for an image to be published to the shared image
    gallery before timing out. If your Packer build is failing on the
    Publishing to Shared Image Gallery step with the error `Original Error:
    context deadline exceeded`, but the image is present when you check your
    Azure dashboard, then you probably need to increase this timeout from
    its default of "60m" (valid time units include `s` for seconds, `m` for
    minutes, and `h` for hours.)
-   `image_publisher` (string) - Image publisher name
    
-   `image_offer` (string) - Image Offer name
-   `image_sku` (string) - Image Sku name
-   `image_version` (string) - Specify a specific version of an OS to boot from.
    Defaults to `latest`. There may be a difference in versions available
    across regions due to image synchronization latency. To ensure a consistent
    version across regions set this value to one that is available in all
    regions where you are deploying.
    
    CLI example
    `az vm image list --location westus --publisher Canonical --offer UbuntuServer --sku 16.04.0-LTS --all`
    
-   `custom_managed_image_resource_group_name` (string) - Specify the source managed image's resource group used to use. If this
    value is set, do not set image\_publisher, image\_offer, image\_sku, or
    image\_version. If this value is set, the value
    `custom_managed_image_name` must also be set. See
    [documentation](https://docs.microsoft.com/en-us/azure/storage/storage-managed-disks-overview#images)
    to learn more about managed images.
    
-   `custom_managed_image_name` (string) - Specify the source managed image's name to use. If this value is set, do
    not set image\_publisher, image\_offer, image\_sku, or image\_version.
    If this value is set, the value
    `custom_managed_image_resource_group_name` must also be set. See
    [documentation](https://docs.microsoft.com/en-us/azure/storage/storage-managed-disks-overview#images)
    to learn more about managed images.
-   `location` (string) - Location
-   `vm_size` (string) - Size of the VM used for building. This can be changed when you deploy a
    VM from your VHD. See
    [pricing](https://azure.microsoft.com/en-us/pricing/details/virtual-machines/)
    information. Defaults to `Standard_A1`.
    
    CLI example `az vm list-sizes --location westus`
    
-   `managed_image_resource_group_name` (string) - Specify the managed image resource group name where the result of the
    Packer build will be saved. The resource group must already exist. If
    this value is set, the value managed_image_name must also be set. See
    documentation to learn more about managed images.
    
-   `managed_image_name` (string) - Specify the managed image name where the result of the Packer build will
    be saved. The image name must not exist ahead of time, and will not be
    overwritten. If this value is set, the value
    managed_image_resource_group_name must also be set. See documentation to
    learn more about managed images.
    
-   `managed_image_storage_account_type` (string) - Specify the storage account
    type for a managed image. Valid values are Standard_LRS and Premium_LRS.
    The default is Standard_LRS.
    
-   `managed_image_os_disk_snapshot_name` (string) - If managed_image_os_disk_snapshot_name is set, a snapshot of the OS disk
    is created with the same name as this value before the VM is captured.
    
-   `managed_image_data_disk_snapshot_prefix` (string) - If
    managed_image_data_disk_snapshot_prefix is set, snapshot of the data
    disk(s) is created with the same prefix as this value before the VM is
    captured.
    
-   `managed_image_zone_resilient` (bool) - Store the image in zone-resilient storage. You need to create it in a
    region that supports [availability
    zones](https://docs.microsoft.com/en-us/azure/availability-zones/az-overview).

-   `plan_info` (PlanInformation) - Plan Info
-   `os_type` (string) - OS type of Virtual Machine. Possible Windows or Linux. 
    
    
-   `lab_virtual_network_name` (string) - Lab Virtual Network Name. If specified, both `lab_subnet_name` also should be specified. If not specified  a default `lab_virtual_network_name` and `lab_subnet_name` will be selected. 
-   `lab_subnet_name` (string) - Lab Subnet Name. If specified, both `lab_virtual_network_name` also should be specified. If not specified  a default `lab_virtual_network_name` and `lab_subnet_name` will be selected. 
-   `lab_name` (string) - Lab Name
-   `lab_resource_group_name` (string) - Resource Group Name where lab is located. 
-   `dtl_artifacts` ([]DtlArtifact) - Dtl Artifacts
-   `vm_name` (string) - VM Name
-   `disk_additional_size` ([]int32) - Additional Disks
    
-   `disk_caching_type` (string) - Disk Caching Type
-   `azure_tags` (map[string]\*string) - the user can define up to 15
    tags. Tag names cannot exceed 512 characters, and tag values cannot exceed
    256 characters. Tags are applied to every resource deployed by a Packer
    build, i.e. Resource Group, VM, NIC, VNET, Public IP, KeyVault, etc.
