import Cookies from "js-cookie";

import { handleErrors } from "./error";

export function fetchMe() {
  return dispatch => {
    dispatch(fetchMeBegin());
    return fetch("/api/users/me", {
      headers: {
        authorization: "Bearer " + Cookies.get("_token")
      }
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(fetchMeSuccess(json));
        return json;
      })
      .catch(error => {
        dispatch(fetchMeFailure(error));
      });
  };
}

export const FETCH_ME_BEGIN = "FETCH_ME_BEGIN";
export const FETCH_ME_SUCCESS = "FETCH_ME_SUCCESS";
export const FETCH_ME_FAILURE = "FETCH_ME_FAILURE";

export const fetchMeBegin = () => ({
  type: FETCH_ME_BEGIN
});

export const fetchMeSuccess = me => ({
  type: FETCH_ME_SUCCESS,
  payload: { me }
});

export const fetchMeFailure = error => ({
  type: FETCH_ME_FAILURE,
  payload: { error }
});
