import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import Home from "./pages/home"
import Signup from "./pages/Signup"

function App() {

  return (
      <Router> 
          <Routes>
            <Route path="/" element={<Home />}/>
            <Route path="/signin" element={<Signup />}/>
          </Routes>
      </Router>
  )
}

export default App
