import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import Home from "./pages/home"
import Signup from "./pages/Signup"
import LogIn from "./pages/Login"
import ProtectedRoute from "./pages/ProtectedRoute"
import PublicRoute from "./pages/PublicRoute"

function App() {

  return (
      <Router> 
          <Routes>
              <Route element={<PublicRoute />}>
                  <Route path="/login" element={<LogIn />} />
                  <Route path="/signup" element={<Signup />} />
              </Route>

              <Route element={<ProtectedRoute />}>
                  <Route path="/" element={<Home />} />
              </Route>
          </Routes>
      </Router>
  )
}

export default App
