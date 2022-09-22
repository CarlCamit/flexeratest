package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUser(t *testing.T) {
	testCases := []struct {
		name     string
		expected *User
	}{
		{
			name: "returns a pointer to a new user struct with desktops and laptops map initialized",
			expected: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  make(map[string]struct{}, 1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := NewUser()
			if diff := cmp.Diff(user, tc.expected); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestAddIfUniqueDesktop(t *testing.T) {
	testCases := []struct {
		name     string
		id       string
		user     *User
		expected *User
	}{
		{
			name: "adds a new entry in the desktops map when given a computer id when the desktops map is empty",
			id:   "123",
			user: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  make(map[string]struct{}, 1),
			},
			expected: &User{
				Desktops: map[string]struct{}{"123": {}},
				Laptops:  make(map[string]struct{}, 1),
			},
		},
		{
			name: "adds a new entry in the desktops map when given a computer id that does not exist",
			id:   "456",
			user: &User{
				Desktops: map[string]struct{}{"123": {}},
				Laptops:  make(map[string]struct{}, 1),
			},
			expected: &User{
				Desktops: map[string]struct{}{"123": {}, "456": {}},
				Laptops:  make(map[string]struct{}, 1),
			},
		},
		{
			name: "ignores add a new entry to the desktops map when given a computer id that does exist",
			id:   "123",
			user: &User{
				Desktops: map[string]struct{}{"123": {}, "456": {}},
				Laptops:  make(map[string]struct{}, 1),
			},
			expected: &User{
				Desktops: map[string]struct{}{"123": {}, "456": {}},
				Laptops:  make(map[string]struct{}, 1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.user.AddIfUniqueDesktop(tc.id)
			if diff := cmp.Diff(tc.user, tc.expected); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestAddIfUniqueLaptop(t *testing.T) {
	testCases := []struct {
		name     string
		id       string
		user     *User
		expected *User
	}{
		{
			name: "adds a new entry in the laptops map when given a computer id when the laptops map is empty",
			id:   "123",
			user: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  make(map[string]struct{}, 1),
			},
			expected: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  map[string]struct{}{"123": {}},
			},
		},
		{
			name: "adds a new entry in the laptops map when given a computer id that does not exist",
			id:   "456",
			user: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  map[string]struct{}{"123": {}},
			},
			expected: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  map[string]struct{}{"123": {}, "456": {}},
			},
		},
		{
			name: "ignores add a new entry to the laptops map when given a computer id that does exist",
			id:   "123",
			user: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  map[string]struct{}{"123": {}, "456": {}},
			},
			expected: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  map[string]struct{}{"123": {}, "456": {}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.user.AddIfUniqueLaptop(tc.id)
			if diff := cmp.Diff(tc.user, tc.expected); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestApplications(t *testing.T) {
	testCases := []struct {
		name     string
		user     *User
		expected int
	}{
		{
			name: "returns a 0 when there are no entries in the desktops or laptops map",
			user: &User{
				Desktops: make(map[string]struct{}, 1),
				Laptops:  make(map[string]struct{}, 1),
			},
			expected: 0,
		},
		{
			name: "returns the length of the desktops map if it is greater than the laptops map",
			user: &User{
				Desktops: map[string]struct{}{"123": {}, "456": {}},
				Laptops:  map[string]struct{}{"789": {}},
			},
			expected: 2,
		},
		{
			name: "returns the length of the laptops map if it is greater than the desktops map",
			user: &User{
				Desktops: map[string]struct{}{"210": {}},
				Laptops:  map[string]struct{}{"123": {}, "456": {}, "789": {}},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.user.Applications() != tc.expected {
				t.Errorf("got a different applications count for a user, got %v, expected %v", tc.user.Applications(), tc.expected)
			}
		})
	}
}

func TestNewUsers(t *testing.T) {
	testCases := []struct {
		name     string
		expected Users
	}{
		{
			name:     "returns a map of Users",
			expected: make(map[string]*User, 1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			users := NewUsers()
			if diff := cmp.Diff(users, tc.expected); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestTotalApplications(t *testing.T) {
	userA := &User{
		Desktops: map[string]struct{}{"123": {}, "456": {}},
		Laptops:  map[string]struct{}{"789": {}},
	}

	userB := &User{
		Desktops: map[string]struct{}{"449": {}, "198": {}, "821": {}, "330": {}},
		Laptops:  map[string]struct{}{"134": {}},
	}

	userC := &User{
		Desktops: make(map[string]struct{}, 1),
		Laptops:  map[string]struct{}{"193": {}},
	}

	userD := &User{
		Desktops: map[string]struct{}{"874": {}, "662": {}},
		Laptops:  map[string]struct{}{"272": {}, "396": {}, "854": {}, "113": {}},
	}

	testCases := []struct {
		name     string
		users    Users
		expected int
	}{
		{
			name:     "returns a 0 if there are no users in the user map",
			users:    NewUsers(),
			expected: 0,
		},
		{
			name: "returns the total number of applications needed for all users",
			users: Users{
				"474": userA,
				"189": userB,
				"551": userC,
				"729": userD,
			},
			expected: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.users.TotalApplications() != tc.expected {
				t.Errorf("got a different total applications count, got %v, expected %v", tc.users.TotalApplications(), tc.expected)
			}
		})
	}
}
