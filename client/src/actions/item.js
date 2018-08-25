import Cookies from "js-cookie";

import { handleErrors } from "./error";

export function fetchItems(collection) {
  return dispatch => {
    dispatch(fetchItemsBegin());
    return fetch("/api/collections/" + collection.uuid + "/items", {
      headers: {
        authorization: "Bearer " + Cookies.get("_token")
      }
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(fetchItemsSuccess(json));
        return json;
      })
      .catch(error => dispatch(fetchItemsFailure(error)));
  };
}

export function addItem(collection, file) {
  return dispatch => {
    dispatch(addItemBegin());
    let formData = new FormData();
    formData.append("file", file);
    return fetch("/api/collections/" + collection.uuid + "/items", {
      method: "POST",
      headers: {
        authorization: "Bearer " + Cookies.get("_token"),
        "X-CSRF-Token": Cookies.get("_csrf")
      },
      body: formData
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(addItemSuccess(json));
        return json;
      })
      .catch(error => dispatch(addItemFailure(error)));
  };
}

export const FETCH_ITEMS_BEGIN = "FETCH_ITEMS_BEGIN";
export const FETCH_ITEMS_SUCCESS = "FETCH_ITEMS_SUCCESS";
export const FETCH_ITEMS_FAILURE = "FETCH_ITEMS_FAILURE";

export const fetchItemsBegin = () => ({
  type: FETCH_ITEMS_BEGIN
});

export const fetchItemsSuccess = items => ({
  type: FETCH_ITEMS_SUCCESS,
  payload: { items }
});

export const fetchItemsFailure = error => ({
  type: FETCH_ITEMS_FAILURE,
  payload: { error }
});

export const ADD_ITEM_BEGIN = "ADD_ITEM_BEGIN";
export const ADD_ITEM_SUCCESS = "ADD_ITEM_SUCCESS";
export const ADD_ITEM_FAILURE = "ADD_ITEM_FAILURE";

export const addItemBegin = () => ({
  type: ADD_ITEM_BEGIN
});

export const addItemSuccess = item => ({
  type: ADD_ITEM_SUCCESS,
  payload: { item }
});

export const addItemFailure = error => ({
  type: ADD_ITEM_FAILURE,
  payload: { error }
});