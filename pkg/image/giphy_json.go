package image

type GiphyResponse struct {
	Data []struct {
		Analytics struct {
			Onclick struct {
				URL string `json:"url"`
			} `json:"onclick"`
			Onload struct {
				URL string `json:"url"`
			} `json:"onload"`
			Onsent struct {
				URL string `json:"url"`
			} `json:"onsent"`
		} `json:"analytics"`
		AnalyticsResponsePayload string `json:"analytics_response_payload"`
		BitlyGifURL              string `json:"bitly_gif_url"`
		BitlyURL                 string `json:"bitly_url"`
		ContentURL               string `json:"content_url"`
		EmbedURL                 string `json:"embed_url"`
		ID                       string `json:"id"`
		Images                   struct {
			_480wStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"480w_still"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			Hd *struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"hd,omitempty"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
		} `json:"images"`
		ImportDatetime   string `json:"import_datetime"`
		IsSticker        int    `json:"is_sticker"`
		Rating           string `json:"rating"`
		Slug             string `json:"slug"`
		Source           string `json:"source"`
		SourcePostURL    string `json:"source_post_url"`
		SourceTld        string `json:"source_tld"`
		Title            string `json:"title"`
		TrendingDatetime string `json:"trending_datetime"`
		Type             string `json:"type"`
		URL              string `json:"url"`
		User             *struct {
			AvatarURL    string `json:"avatar_url"`
			BannerImage  string `json:"banner_image"`
			BannerURL    string `json:"banner_url"`
			Description  string `json:"description"`
			DisplayName  string `json:"display_name"`
			InstagramURL string `json:"instagram_url"`
			IsVerified   bool   `json:"is_verified"`
			ProfileURL   string `json:"profile_url"`
			Username     string `json:"username"`
			WebsiteURL   string `json:"website_url"`
		} `json:"user,omitempty"`
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
		Status     int    `json:"status"`
	} `json:"meta"`
	Pagination struct {
		Count      int `json:"count"`
		Offset     int `json:"offset"`
		TotalCount int `json:"total_count"`
	} `json:"pagination"`
}
