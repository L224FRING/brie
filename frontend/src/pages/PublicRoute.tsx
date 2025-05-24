// PublicRoute.tsx
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

const PublicRoute = () => {
  const { authenticated, isLoading } = useAuth();

  if (isLoading) return <div>Loading...</div>;

  return !authenticated ? <Outlet /> : <Navigate to="/" replace />;
};

export default PublicRoute;

