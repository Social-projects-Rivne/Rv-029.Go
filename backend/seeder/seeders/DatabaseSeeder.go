package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/src"
)

func Run() {
	seeders.Call(UsersTableSeeder{})
}
