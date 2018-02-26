// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func BackendResource() *schema.Resource {
	return &schema.Resource{
		Create: createBackend,
		Read:   readBackend,
		Update: updateBackend,
		Delete: deleteBackend,
		Schema: map[string]*schema.Schema{
			// Required
			"backendset_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"drain": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"offline": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updateBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type BackendResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Backend
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *BackendResourceCrud) buildID() string {
	return s.D.Get("ip_address").(string) + ":" + strconv.Itoa(s.D.Get("port").(int))
}

func (s *BackendResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		// Always inferred this way
		return s.buildID()
	}
	return ""
}

func (s *BackendResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendResourceCrud) Create() error {
	request := oci_load_balancer.CreateBackendRequest{}

	if backendsetName, ok := s.D.GetOk("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if backup, ok := s.D.GetOk("backup"); ok {
		tmp := backup.(bool)
		request.Backup = &tmp
	}

	if drain, ok := s.D.GetOk("drain"); ok {
		tmp := drain.(bool)
		request.Drain = &tmp
	}

	if ipAddress, ok := s.D.GetOk("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if loadBalancerId, ok := s.D.GetOk("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if offline, ok := s.D.GetOk("offline"); ok {
		tmp := offline.(bool)
		request.Offline = &tmp
	}

	if port, ok := s.D.GetOk("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if weight, ok := s.D.GetOk("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	response, err := s.Client.CreateBackend(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *BackendResourceCrud) Get() error {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetBackendRequest{}

	tmp := s.buildID()
	request.BackendName = &tmp

	if backendSetName, ok := s.D.GetOk("backendset_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOk("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	response, err := s.Client.GetBackend(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}

	s.Res = &response.Backend
	return nil
}

func (s *BackendResourceCrud) Update() error {
	request := oci_load_balancer.UpdateBackendRequest{}

	if backendName, ok := s.D.GetOk("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendsetName, ok := s.D.GetOk("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if backup, ok := s.D.GetOk("backup"); ok {
		tmp := backup.(bool)
		request.Backup = &tmp
	}

	if drain, ok := s.D.GetOk("drain"); ok {
		tmp := drain.(bool)
		request.Drain = &tmp
	}

	if loadBalancerId, ok := s.D.GetOk("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if offline, ok := s.D.GetOk("offline"); ok {
		tmp := offline.(bool)
		request.Offline = &tmp
	}

	if weight, ok := s.D.GetOk("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	response, err := s.Client.UpdateBackend(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *BackendResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteBackendRequest{}

	if backendName, ok := s.D.GetOk("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendSetName, ok := s.D.GetOk("backendset_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOk("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	response, err := s.Client.DeleteBackend(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	s.D.SetId(*workReqID)
	return nil
}

func (s *BackendResourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	if s.Res.Backup != nil {
		s.D.Set("backup", *s.Res.Backup)
	}

	if s.Res.Drain != nil {
		s.D.Set("drain", *s.Res.Drain)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Offline != nil {
		s.D.Set("offline", *s.Res.Offline)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.Weight != nil {
		s.D.Set("weight", *s.Res.Weight)
	}

}