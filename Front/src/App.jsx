import { Routes, Route, Link } from 'react-router-dom'
import Flcon from './FIcon'
import './App.css'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Flcon />} />

    </Routes>
  )
}

export default App