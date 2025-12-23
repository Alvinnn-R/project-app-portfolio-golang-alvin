package dto

import (
	"testing"
)

func TestProfileRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ProfileRequest
		wantErr bool
	}{
		{
			name: "valid profile request",
			req: ProfileRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			req: ProfileRequest{
				Name:  "",
				Email: "john@example.com",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			req: ProfileRequest{
				Name:  "John Doe",
				Email: "",
			},
			wantErr: true,
		},
		{
			name: "name too long",
			req: ProfileRequest{
				Name:  string(make([]byte, 101)),
				Email: "john@example.com",
			},
			wantErr: true,
		},
		{
			name: "email too long",
			req: ProfileRequest{
				Name:  "John Doe",
				Email: string(make([]byte, 101)),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExperienceRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ExperienceRequest
		wantErr bool
	}{
		{
			name: "valid experience request",
			req: ExperienceRequest{
				Title:        "Software Developer",
				Organization: "Tech Corp",
				Type:         "work",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			req: ExperienceRequest{
				Title:        "",
				Organization: "Tech Corp",
				Type:         "work",
			},
			wantErr: true,
		},
		{
			name: "empty organization",
			req: ExperienceRequest{
				Title:        "Software Developer",
				Organization: "",
				Type:         "work",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			req: ExperienceRequest{
				Title:        "Software Developer",
				Organization: "Tech Corp",
				Type:         "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid internship type",
			req: ExperienceRequest{
				Title:        "Intern",
				Organization: "Tech Corp",
				Type:         "internship",
			},
			wantErr: false,
		},
		{
			name: "valid campus type",
			req: ExperienceRequest{
				Title:        "Research Assistant",
				Organization: "University",
				Type:         "campus",
			},
			wantErr: false,
		},
		{
			name: "valid competition type",
			req: ExperienceRequest{
				Title:        "Hackathon",
				Organization: "Tech Corp",
				Type:         "competition",
			},
			wantErr: false,
		},
		{
			name: "title too long",
			req: ExperienceRequest{
				Title:        string(make([]byte, 201)),
				Organization: "Tech Corp",
				Type:         "work",
			},
			wantErr: true,
		},
		{
			name: "organization too long",
			req: ExperienceRequest{
				Title:        "Developer",
				Organization: string(make([]byte, 201)),
				Type:         "work",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperienceRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSkillRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     SkillRequest
		wantErr bool
	}{
		{
			name: "valid skill request",
			req: SkillRequest{
				Category: "Programming",
				Name:     "Go",
				Level:    "intermediate",
			},
			wantErr: false,
		},
		{
			name: "empty category",
			req: SkillRequest{
				Category: "",
				Name:     "Go",
				Level:    "intermediate",
			},
			wantErr: true,
		},
		{
			name: "empty name",
			req: SkillRequest{
				Category: "Programming",
				Name:     "",
				Level:    "intermediate",
			},
			wantErr: true,
		},
		{
			name: "invalid level",
			req: SkillRequest{
				Category: "Programming",
				Name:     "Go",
				Level:    "expert",
			},
			wantErr: true,
		},
		{
			name: "valid beginner level",
			req: SkillRequest{
				Category: "Programming",
				Name:     "Go",
				Level:    "beginner",
			},
			wantErr: false,
		},
		{
			name: "valid advanced level",
			req: SkillRequest{
				Category: "Programming",
				Name:     "Go",
				Level:    "advanced",
			},
			wantErr: false,
		},
		{
			name: "empty level is valid",
			req: SkillRequest{
				Category: "Programming",
				Name:     "Go",
				Level:    "",
			},
			wantErr: false,
		},
		{
			name: "category too long",
			req: SkillRequest{
				Category: string(make([]byte, 101)),
				Name:     "Go",
				Level:    "intermediate",
			},
			wantErr: true,
		},
		{
			name: "name too long",
			req: SkillRequest{
				Category: "Programming",
				Name:     string(make([]byte, 101)),
				Level:    "intermediate",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("SkillRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ProjectRequest
		wantErr bool
	}{
		{
			name: "valid project request",
			req: ProjectRequest{
				Title:       "My Project",
				Description: "A cool project",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			req: ProjectRequest{
				Title:       "",
				Description: "A cool project",
			},
			wantErr: true,
		},
		{
			name: "title too long",
			req: ProjectRequest{
				Title:       string(make([]byte, 201)),
				Description: "A cool project",
			},
			wantErr: true,
		},
		{
			name: "empty description is valid",
			req: ProjectRequest{
				Title:       "My Project",
				Description: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPublicationRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     PublicationRequest
		wantErr bool
	}{
		{
			name: "valid publication request",
			req: PublicationRequest{
				Title: "My Publication",
				Year:  2024,
			},
			wantErr: false,
		},
		{
			name: "empty title",
			req: PublicationRequest{
				Title: "",
				Year:  2024,
			},
			wantErr: true,
		},
		{
			name: "year too old",
			req: PublicationRequest{
				Title: "My Publication",
				Year:  1899,
			},
			wantErr: true,
		},
		{
			name: "year in future",
			req: PublicationRequest{
				Title: "My Publication",
				Year:  2101,
			},
			wantErr: true,
		},
		{
			name: "title too long",
			req: PublicationRequest{
				Title: string(make([]byte, 201)),
				Year:  2024,
			},
			wantErr: true,
		},
		{
			name: "valid minimum year",
			req: PublicationRequest{
				Title: "My Publication",
				Year:  1900,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicationRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContactRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     ContactRequest
		wantErr bool
	}{
		{
			name: "valid contact request",
			req: ContactRequest{
				Name:    "John Doe",
				Email:   "john@example.com",
				Subject: "Hello",
				Message: "This is a message",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			req: ContactRequest{
				Name:    "",
				Email:   "john@example.com",
				Subject: "Hello",
				Message: "This is a message",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			req: ContactRequest{
				Name:    "John Doe",
				Email:   "",
				Subject: "Hello",
				Message: "This is a message",
			},
			wantErr: true,
		},
		{
			name: "empty message",
			req: ContactRequest{
				Name:    "John Doe",
				Email:   "john@example.com",
				Subject: "Hello",
				Message: "",
			},
			wantErr: true,
		},
		{
			name: "empty subject is valid",
			req: ContactRequest{
				Name:    "John Doe",
				Email:   "john@example.com",
				Subject: "",
				Message: "This is a message",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContactRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
