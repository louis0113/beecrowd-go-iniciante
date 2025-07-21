package main

import "fmt"

func main (){
  const maxMin = 60
  const maxHour = 24
  const m3 = "MINUTO(S)" 
  const m2 = "HORA(S)"
  const m1 = "O JOGO DUROU"
  var hi,mi,hf,mf,ha,ma int 

  fmt.Scanf("%d %d %d %d\n", &hi, &mi, &hf, &mf)

  if hi < hf && mi == mf{
    ha = hf - hi 
    ma = 0
  } else if hi > hf && mi == mf{
    ha = (maxHour - hi) + hf 
    ma = 0
  } else if hi == hf && mi < mf{
    ha = 0
    ma = mf - mi
  } else if hi == hf && mi > mf {
    ha = 0
    ma = (maxMin - mi) + mf 
  } else if hi < hf && mi < mf{
    ha = hf - hi
    ma = mf - mi
  } else if hi > hf && mi > mf{
    ha = (maxHour - hi) + hf 
    ma = (maxMin - mi) + mf
  } else if hi < hf && mi > mf && (hf - hi) == 1{
    ha = (hf - hi) - 1
    ma = (maxMin - mi) + mf
  } else if hi > hf && mi < mf {
    ha = (maxHour - hi) + hf 
    ma = mf - mi
  } else if hi == hf && mi == hf {
    ha = 24
    ma = 0
  } else if hi < hf && mi > mf && (hf - hi) != 1{

    ha = (hf - hi) - 1
    ma = (maxMin - mi) + mf
  }

  fmt.Printf("%s %d %s E %d %s\n", m1,ha,m2,ma,m3)
}
