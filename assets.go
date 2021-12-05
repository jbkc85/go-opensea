package opensea

type AssetsResponse struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	ID                   int64         `json:"id"`
	TokenID              string        `json:"token_id"`
	NumSales             int           `json:"num_sales"`
	BackgroundColor      string        `json:"background_color"`
	ImageURL             string        `json:"image_url"`
	ImagePreviewURL      string        `json:"image_preview_url"`
	ImageThumbnailURL    string        `json:"image_thumbnail_url"`
	AnimationURL         string        `json:"animation_url"`
	AnimationOriginalArt string        `json:"animation_original_art"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	ExternalLink         string        `json:"external_link"`
	Contract             AssetContract `json:"asset_contract"`
	Permalink            string        `json:"permalink"`
	Traits               []AssetTraits `json:"traits"`
}

type AssetContract struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"` // this should be a separate struct
	CreatedDate                 string `json:"created_date"`
	Name                        string `json:"name"`
	NFTVersion                  string `jsion:"nft_version"`
	OpenseaVersion              string `json:"opensea_version"`
	Owner                       int64  `json:"owner"`
	SchemaName                  string `json:"schema_name"` // this should be a separate struct
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
	Description                 string `json:"description"`
	ExternalLink                string `json:"external_link"`
	ImageURL                    string `json:"image_url"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address"`
}

type AssetTraits struct {
	Type        string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type"`
	MaxValue    int64       `jsson:"max_value"`
	Count       int64       `json:"trait_count"`
	Order       string      `json:"order"`
}
