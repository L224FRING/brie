// ProtectedRoute.tsx
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
const ProtectedRoute = () => {
  const { authenticated, isLoading } = useAuth();

  if (isLoading) return <div>Loading...</div>;

  return authenticated ? <Outlet /> : <Navigate to="/login" replace />;
};

export default ProtectedRoute;

