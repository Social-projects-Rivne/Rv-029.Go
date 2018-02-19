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
<<<<<<< HEAD
=======
	seeders.Call(IssueTableSeeder{})
>>>>>>> 3b899faad40f9803ee884c54936abe6a1354b292
}
