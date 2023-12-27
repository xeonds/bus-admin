import {jwtDecode} from "jwt-decode";

export default function isLogin() {
  const token = localStorage.getItem("token");
  if (!token || token == "" || token == "undefined") return false;
  if (Date.now() > jwtDecode(localStorage.getItem("token")!).exp!)
    return false;
  return true;
}
