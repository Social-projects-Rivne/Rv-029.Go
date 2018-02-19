package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/src"
)

//Run .
func Run() {
	seeders.Call(UsersTableSeeder{})
	seeders.Call(ProjectTableSeeder{})
	seeders.Call(BoardTableSeeder{})
	seeders.Call(SprintTableSeeder{})
	seeders.Call(IssueTableSeeder{})
}
