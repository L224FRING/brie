import { useState, ChangeEvent, FormEvent } from "react";
import "./Signup.css";
import { Link, useNavigate } from "react-router-dom";

interface FormData {
  username: string;
  password: string;
}

const Signup = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState<FormData>({
    username: "",
    password: "",
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
      setFormData({
          ...formData,
          [e.target.name]: e.target.value,
      });
  };
  const handleSubmit = async (e: FormEvent) => {
      e.preventDefault();

      try {
          const response = await fetch("http://localhost/api/auth/sign-in", {
              method: "POST",
              headers: {
                  "Content-Type": "application/json",
              },
              credentials: 'include',
              body: JSON.stringify(formData),
          });

          if (!response.ok) {
              const errorData = await response.json();
              alert(`Signup failed: ${errorData.message || response.statusText}`);
              return;
          }

          alert("Signup Successful!");
          navigate("/")
          // Optionally, reset the form or redirect the user
      } catch (error) {
          console.error("Error during signup:", error);
          alert("An error occurred during signup. Please try again later.");
      }
  };


  return (
    <div className="signup-container">
      <form className="signup-form" onSubmit={handleSubmit}>
        <h2>Sign Up</h2>

        <label>Username</label>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleChange}
          required
        />

        <label>Password</label>
        <input
          type="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          required
        />

        <button type="submit">Sign Up</button> 
        <div className="link">
            already have and account <Link to="/login">Log In</Link>
        </div>
      </form>
    </div>
  );
};

export default Signup;


