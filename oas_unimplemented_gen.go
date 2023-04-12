// Code generated by ogen, DO NOT EDIT.

package client

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateAdType implements createAdType operation.
//
// POST /v1/adtypes
func (UnimplementedHandler) CreateAdType(ctx context.Context, req OptCreateAdTypeReq) (r *AdType, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChannel implements createChannel operation.
//
// POST /v1/channel
func (UnimplementedHandler) CreateChannel(ctx context.Context, req OptCreateChannelReq) (r *Channel, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChannelSiteMap implements createChannelSiteMap operation.
//
// POST /v1/channelSite
func (UnimplementedHandler) CreateChannelSiteMap(ctx context.Context, req OptCreateChannelSiteMapReq) (r *ChannelSiteMap, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateForChannelAdType implements createForChannelAdType operation.
//
// POST /v1/channel/{channelId}/adtypes
func (UnimplementedHandler) CreateForChannelAdType(ctx context.Context, req OptCreateForChannelAdTypeReq, params CreateForChannelAdTypeParams) (r *AdType, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSite implements createSite operation.
//
// POST /v1/site
func (UnimplementedHandler) CreateSite(ctx context.Context, req OptCreateSiteReq) (r *Site, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateZone implements createZone operation.
//
// POST /v1/zone
func (UnimplementedHandler) CreateZone(ctx context.Context, req OptCreateZoneReq) (r *Zone, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteAdType implements deleteAdType operation.
//
// GET /v1/adtypes/{id}/delete
func (UnimplementedHandler) DeleteAdType(ctx context.Context, params DeleteAdTypeParams) error {
	return ht.ErrNotImplemented
}

// DeleteChannel implements deleteChannel operation.
//
// GET /v1/channel/{id}/delete
func (UnimplementedHandler) DeleteChannel(ctx context.Context, params DeleteChannelParams) error {
	return ht.ErrNotImplemented
}

// DeleteChannelSiteMap implements deleteChannelSiteMap operation.
//
// GET /v1/channel/{channelId}/site/{siteId}/delete
func (UnimplementedHandler) DeleteChannelSiteMap(ctx context.Context, params DeleteChannelSiteMapParams) error {
	return ht.ErrNotImplemented
}

// DeleteForChannelAdType implements deleteForChannelAdType operation.
//
// GET /v1/channel/{channelId}/adtypes/{id}/delete
func (UnimplementedHandler) DeleteForChannelAdType(ctx context.Context, params DeleteForChannelAdTypeParams) error {
	return ht.ErrNotImplemented
}

// FilterSite implements filterSite operation.
//
// GET /v1/fast/site
func (UnimplementedHandler) FilterSite(ctx context.Context, params FilterSiteParams) (r FilterSiteOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChannel implements getChannel operation.
//
// GET /v1/channel/{id}
func (UnimplementedHandler) GetChannel(ctx context.Context, params GetChannelParams) (r *Channel, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChannelSiteMap implements getChannelSiteMap operation.
//
// GET /v1/channel/{channelId}/site/{siteId}
func (UnimplementedHandler) GetChannelSiteMap(ctx context.Context, params GetChannelSiteMapParams) (r *ChannelSiteMap, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPrioritiesChannel implements getPrioritiesChannel operation.
//
// GET /v1/channel/{id}/priorities
func (UnimplementedHandler) GetPrioritiesChannel(ctx context.Context, params GetPrioritiesChannelParams) (r []Priority, _ error) {
	return r, ht.ErrNotImplemented
}

// GetSite implements getSite operation.
//
// GET /v1/site/{id}
func (UnimplementedHandler) GetSite(ctx context.Context, params GetSiteParams) (r *Site, _ error) {
	return r, ht.ErrNotImplemented
}

// GetZone implements getZone operation.
//
// GET /v1/zone/{id}
func (UnimplementedHandler) GetZone(ctx context.Context, params GetZoneParams) (r *Zone, _ error) {
	return r, ht.ErrNotImplemented
}

// ListAdTypes implements listAdTypes operation.
//
// GET /v1/adtypes
func (UnimplementedHandler) ListAdTypes(ctx context.Context, params ListAdTypesParams) (r *AdTypeList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListChannelSiteMaps implements listChannelSiteMaps operation.
//
// GET /v1/channelSite
func (UnimplementedHandler) ListChannelSiteMaps(ctx context.Context, params ListChannelSiteMapsParams) (r *ChannelSiteMapList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListChannels implements listChannels operation.
//
// GET /v1/channel
func (UnimplementedHandler) ListChannels(ctx context.Context, params ListChannelsParams) (r *ChannelList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListChannelsForSiteChannelSiteMap implements listChannelsForSiteChannelSiteMap operation.
//
// GET /v1/channelsInSite/{siteId}
func (UnimplementedHandler) ListChannelsForSiteChannelSiteMap(ctx context.Context, params ListChannelsForSiteChannelSiteMapParams) (r *ChannelList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListForChannelAdType implements listForChannelAdType operation.
//
// GET /v1/channel/{channelId}/adtypes
func (UnimplementedHandler) ListForChannelAdType(ctx context.Context, params ListForChannelAdTypeParams) (r *AdTypeList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSites implements listSites operation.
//
// GET /v1/site
func (UnimplementedHandler) ListSites(ctx context.Context, params ListSitesParams) (r *SiteList, _ error) {
	return r, ht.ErrNotImplemented
}

// ListZones implements listZones operation.
//
// GET /v1/zone
func (UnimplementedHandler) ListZones(ctx context.Context, params ListZonesParams) (r *ZoneList, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateChannel implements updateChannel operation.
//
// PUT /v1/channel/{id}
func (UnimplementedHandler) UpdateChannel(ctx context.Context, req OptUpdateChannelReq, params UpdateChannelParams) (r *Channel, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateChannelSiteMap implements updateChannelSiteMap operation.
//
// PUT /v1/channelSite
func (UnimplementedHandler) UpdateChannelSiteMap(ctx context.Context, req OptUpdateChannelSiteMapReq) (r *ChannelSiteMap, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateSite implements updateSite operation.
//
// PUT /v1/site/{id}
func (UnimplementedHandler) UpdateSite(ctx context.Context, req OptUpdateSiteReq, params UpdateSiteParams) (r *Site, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateZone implements updateZone operation.
//
// PUT /v1/zone/{id}
func (UnimplementedHandler) UpdateZone(ctx context.Context, req OptUpdateZoneReq, params UpdateZoneParams) (r *Zone, _ error) {
	return r, ht.ErrNotImplemented
}
