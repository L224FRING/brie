import { useNavigate } from "react-router-dom";
const Home = () => {
  const navigate = useNavigate();
  const handleLogout = async () => {
    try {
      await fetch("http://localhost/api/auth/log-out", {
        method: "POST",
        credentials: "include", // ensures the cookie is sent
      });
      navigate("/login");
    } catch (error) {
      console.error("Logout failed", error);
    }
  };
  return (
  <div>
        <h1>Home Page</h1>;
        <button onClick={handleLogout}>Logout</button>
    </div>
  )
};
export default Home;


