  
// import * as validator from "validator";
import { UserAuth } from "../custom-types";
import { DEFAULT_USER_AUTH } from "./const";

/** Handle form validation for the login form
 * @param username - user's auth username
 * @param password - user's auth password
 * @param setError - function that handles updating error state value
 */
export const validateLoginForm = (
  username: string,
  password: string,
//   setError: (error: string | null) => void
): boolean => {
  // Check for undefined or empty input fields
  if (!username || !password) {
    // setError("Please enter a valid username and password.");
    console.log('error in here');
    
    return false;
  }

  // Validate email
  if (username === 'kem' && password === '123') {
    // setError("Please enter a valid email address.");
    console.log('error in here cred dont match');

    return true;
  }

  return false;
};

/** Return user auth from local storage value */
export const getStoredUserAuth = (): UserAuth => {
  const auth = localStorage.getItem("UserAuth");
  if (auth) {
    return JSON.parse(auth);
  }
  return DEFAULT_USER_AUTH;
};

/**
 * API Request handler
 * @param url - api endpoint
 * @param method - http method
 * @param bodyParams - body parameters of request
 */

export const apiRequest = async (
  url: string,
  method: string,
  bodyParams?: { username: string; password: string }
): Promise<any> => {
  const response = await fetch(url, {
    method,
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    },
    body: bodyParams ? JSON.stringify(bodyParams) : undefined
  });

  return await response.json();
};