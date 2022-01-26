package api

import "time"

type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	BlurHash       string `json:"blur_hash"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Sponsorship            struct {
		ImpressionUrls []string `json:"impression_urls"`
		Tagline        string   `json:"tagline"`
		TaglineURL     string   `json:"tagline_url"`
		Sponsor        struct {
			ID              string      `json:"id"`
			UpdatedAt       string      `json:"updated_at"`
			Username        string      `json:"username"`
			Name            string      `json:"name"`
			FirstName       string      `json:"first_name"`
			LastName        interface{} `json:"last_name"`
			TwitterUsername string      `json:"twitter_username"`
			PortfolioURL    string      `json:"portfolio_url"`
			Bio             string      `json:"bio"`
			Location        interface{} `json:"location"`
			Links           struct {
				Self      string `json:"self"`
				HTML      string `json:"html"`
				Photos    string `json:"photos"`
				Likes     string `json:"likes"`
				Portfolio string `json:"portfolio"`
				Following string `json:"following"`
				Followers string `json:"followers"`
			} `json:"links"`
			ProfileImage struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Large  string `json:"large"`
			} `json:"profile_image"`
			InstagramUsername string `json:"instagram_username"`
			TotalCollections  int    `json:"total_collections"`
			TotalLikes        int    `json:"total_likes"`
			TotalPhotos       int    `json:"total_photos"`
			AcceptedTos       bool   `json:"accepted_tos"`
			ForHire           bool   `json:"for_hire"`
			Social            struct {
				InstagramUsername string      `json:"instagram_username"`
				PortfolioURL      string      `json:"portfolio_url"`
				TwitterUsername   string      `json:"twitter_username"`
				PaypalEmail       interface{} `json:"paypal_email"`
			} `json:"social"`
		} `json:"sponsor"`
	} `json:"sponsorship"`
}

type PaginatedPhoto struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}

type Topic struct {
	ID                          string        `json:"id"`
	Slug                        string        `json:"slug"`
	Title                       string        `json:"title"`
	Description                 string        `json:"description"`
	PublishedAt                 string        `json:"published_at"`
	UpdatedAt                   string        `json:"updated_at"`
	StartsAt                    time.Time     `json:"starts_at"`
	EndsAt                      time.Time     `json:"ends_at"`
	OnlySubmissionsAfter        interface{}   `json:"only_submissions_after"`
	Featured                    bool          `json:"featured"`
	TotalPhotos                 int           `json:"total_photos"`
	CurrentUserContributions    []interface{} `json:"current_user_contributions"`
	TotalCurrentUserSubmissions interface{}   `json:"total_current_user_submissions"`
	Links                       struct {
		Self   string `json:"self"`
		HTML   string `json:"html"`
		Photos string `json:"photos"`
	} `json:"links"`
	Status string `json:"status"`
	Owners []struct {
		ID              string      `json:"id"`
		UpdatedAt       string      `json:"updated_at"`
		Username        string      `json:"username"`
		Name            string      `json:"name"`
		FirstName       string      `json:"first_name"`
		LastName        interface{} `json:"last_name"`
		TwitterUsername string      `json:"twitter_username"`
		PortfolioURL    string      `json:"portfolio_url"`
		Bio             string      `json:"bio"`
		Location        interface{} `json:"location"`
		Links           struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
		ProfileImage struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		InstagramUsername string `json:"instagram_username"`
		TotalCollections  int    `json:"total_collections"`
		TotalLikes        int    `json:"total_likes"`
		TotalPhotos       int    `json:"total_photos"`
		AcceptedTos       bool   `json:"accepted_tos"`
		ForHire           bool   `json:"for_hire"`
		Social            struct {
			InstagramUsername string      `json:"instagram_username"`
			PortfolioURL      string      `json:"portfolio_url"`
			TwitterUsername   string      `json:"twitter_username"`
			PaypalEmail       interface{} `json:"paypal_email"`
		} `json:"social"`
	} `json:"owners"`
	CoverPhoto struct {
		ID             string `json:"id"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		PromotedAt     string `json:"promoted_at"`
		Width          int    `json:"width"`
		Height         int    `json:"height"`
		Color          string `json:"color"`
		BlurHash       string `json:"blur_hash"`
		Description    string `json:"description"`
		AltDescription string `json:"alt_description"`
		Urls           struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
		Links struct {
			Self             string `json:"self"`
			HTML             string `json:"html"`
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
		} `json:"links"`
		Categories             []interface{} `json:"categories"`
		Likes                  int           `json:"likes"`
		LikedByUser            bool          `json:"liked_by_user"`
		CurrentUserCollections []interface{} `json:"current_user_collections"`
		Sponsorship            interface{}   `json:"sponsorship"`
		TopicSubmissions       struct {
			Spirituality struct {
				Status     string `json:"status"`
				ApprovedOn string `json:"approved_on"`
			} `json:"spirituality"`
			Holidays struct {
				Status     string `json:"status"`
				ApprovedOn string `json:"approved_on"`
			} `json:"holidays"`
		} `json:"topic_submissions"`
		User struct {
			ID              string `json:"id"`
			UpdatedAt       string `json:"updated_at"`
			Username        string `json:"username"`
			Name            string `json:"name"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			TwitterUsername string `json:"twitter_username"`
			PortfolioURL    string `json:"portfolio_url"`
			Bio             string `json:"bio"`
			Location        string `json:"location"`
			Links           struct {
				Self      string `json:"self"`
				HTML      string `json:"html"`
				Photos    string `json:"photos"`
				Likes     string `json:"likes"`
				Portfolio string `json:"portfolio"`
				Following string `json:"following"`
				Followers string `json:"followers"`
			} `json:"links"`
			ProfileImage struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Large  string `json:"large"`
			} `json:"profile_image"`
			InstagramUsername string `json:"instagram_username"`
			TotalCollections  int    `json:"total_collections"`
			TotalLikes        int    `json:"total_likes"`
			TotalPhotos       int    `json:"total_photos"`
			AcceptedTos       bool   `json:"accepted_tos"`
			ForHire           bool   `json:"for_hire"`
			Social            struct {
				InstagramUsername string      `json:"instagram_username"`
				PortfolioURL      string      `json:"portfolio_url"`
				TwitterUsername   string      `json:"twitter_username"`
				PaypalEmail       interface{} `json:"paypal_email"`
			} `json:"social"`
		} `json:"user"`
	} `json:"cover_photo"`
	PreviewPhotos []struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		BlurHash  string `json:"blur_hash"`
		Urls      struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
	} `json:"preview_photos"`
}

type PaginatedTopic struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Topic `json:"results"`
}

type PaginatedTopicPhotos struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}
