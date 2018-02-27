package seeder

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/seeder/src"
	"math/rand"
	"time"
)

func Run() {
	seeders.Call(ProjectTableSeeder{})
	seeders.Call(UsersTableSeeder{})
	seeders.Call(BoardTableSeeder{})
	seeders.Call(SprintTableSeeder{})
	seeders.Call(IssueTableSeeder{})
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randomFromArray(array []string) string {
	return array[rand.Intn(len(array))]
}
