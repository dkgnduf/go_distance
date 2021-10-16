package main

import(
	"fmt"
	"math"
	"time"
)


type Coordinates struct
{
	Longitude float64
	Latitude float64
}
// 경위도 구조체

func degree_to_radius(degree float64) float64 {
	return degree * math.Pi / 180
}
// 각도 -> 라디

func (origin Coordinates) distance(destination Coordinates) float64{
	var dLat, dLon, a, b float64
	const R = 6371 //지구반경
	dLat = degree_to_radius(destination.Latitude - origin.Latitude)
	dLon = degree_to_radius(destination.Longitude - origin.Longitude)
	a = (math.Sin(dLat/2)*math.Sin(dLat/2) +
 		math.Cos(degree_to_radius(origin.Latitude))*
 			math.Cos(degree_to_radius(destination.Latitude))*math.Sin(dLon/2)*
 			math.Sin(dLon/2))
	b = 2*math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * b * 1000
}
// 하버사인 공식을 이용한 두 점 사이의 거리 찾기

func main(){
	coordA := Coordinates{}
	coordB := Coordinates{} // A,B 구조체 선언
	var n1,n2,n3,n4 float64
	var input_num, input_num2 int //n1~n4, 입력 숫자 변수 선언

	fmt.Println("Input Original : Longitude,Latitude")
	input_num, _ = fmt.Scanf("%f,%f\n", &n1,&n2)

	for {
		if input_num != 2 {
			fmt.Println("float 2 input Please")
			input_num, _ = fmt.Scanf("%f,%f\n", &n1, &n2)
			continue
		}else{
			if n1 < -180 || n1 > 180 || n2 < -90 || n2 > 90{
				fmt.Println("-180<=Longitude<=180, -90<=Latitude<=90 \n")
				fmt.Scanf("%f,%f\n", &n1,&n2)
			}else{
				break
			}
		}
	} //n1,n2 입력 2개가 아니거나, 경위도 범위에 벗어나는 경우 예외처리

	fmt.Println("Input Destination : Longitude,Latitude")
	input_num2, _ = fmt.Scanf("%f,%f\n", &n3,&n4)

	for {
		if input_num2 != 2 {
			fmt.Println("float 2 input Please")
			input_num2, _ = fmt.Scanf("%f,%f\n", &n3, &n4)
			continue
		}else{
			if n3 < -180 || n3 > 180 || n4 < -90 || n4 > 90{
				fmt.Println("-180<=Longitude<=180, -90<=Latitude<=90 \n")
				fmt.Scanf("%f,%f\n", &n3,&n4)
			}else{
				break
			}
		}
	} //n3,n4 입력 2개가 아니거나, 경위도 범위에 벗어나는 경우 예외처리

	coordA.Longitude = n1
	coordA.Latitude = n2
	coordB.Longitude = n3
	coordB.Latitude = n4

 	var distance = coordA.distance(coordB)
 	fmt.Printf("distance between two points %f meters.\n", distance) //n1~n4 입력, 거리 계산
	time.Sleep(8 * time.Second)
}
