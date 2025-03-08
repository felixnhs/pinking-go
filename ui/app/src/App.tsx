import './App.css'
import { Page } from './components/Page/index.jsx'
import { picsArr } from './utils/pics.js'
import { Swiper } from './components/Swiper.jsx'

export function App() {

  function easeInOut(x: number) {
    "main thread"
    return x < 0.5
      ? 2 * x * x
      : 1 - Math.pow(-2 * x + 2, 2) / 2;
  }

  return (
    <Page>
      <Swiper data={picsArr} main-thread:easing={easeInOut} itemWidth={390} />
    </Page>
  )
}
