import Cookies from 'js-cookie';
import { handleErrors } from './error';

export function fetchCollections() {
  return dispatch => {
    dispatch(fetchCollectionsBegin());
    return fetch('/api/collections', {
      headers: {
        authorization: 'Bearer ' + Cookies.get('_token')
      }
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(fetchCollectionsSuccess(json));
        return json;
      })
      .catch(error => dispatch(fetchCollectionsFailure(error)));
  };
}

export function fetchCollection(slug) {
  return dispatch => {
    dispatch(fetchCollectionBegin());
    return fetch('/api/collections/' + slug, {
      headers: {
        authorization: 'Bearer ' + Cookies.get('_token')
      }
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(fetchCollectionSuccess(json));
        return json;
      })
      .catch(error => dispatch(fetchCollectionFailure(error)));
  };
}

export function addCollection(name) {
  return dispatch => {
    dispatch(addCollectionBegin());
    return fetch('/api/collections', {
      method: 'POST',
      headers: {
        authorization: 'Bearer ' + Cookies.get('_token'),
        'X-CSRF-Token': Cookies.get('_csrf'),
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ name })
    })
      .then(handleErrors)
      .then(res => res.json())
      .then(json => {
        dispatch(addCollectionSuccess(json));
        return json;
      })
      .catch(error => dispatch(addCollectionFailure(error)));
  };
}

export const FETCH_COLLECTIONS_BEGIN = 'FETCH_COLLECTIONS_BEGIN';
export const FETCH_COLLECTIONS_SUCCESS = 'FETCH_COLLECTIONS_SUCCESS';
export const FETCH_COLLECTIONS_FAILURE = 'FETCH_COLLECTIONS_FAILURE';

export const fetchCollectionsBegin = () => ({
  type: FETCH_COLLECTIONS_BEGIN
});

export const fetchCollectionsSuccess = collections => ({
  type: FETCH_COLLECTIONS_SUCCESS,
  payload: { collections }
});

export const fetchCollectionsFailure = error => ({
  type: FETCH_COLLECTIONS_FAILURE,
  payload: { error }
});

export const FETCH_COLLECTION_BEGIN = 'FETCH_COLLECTION_BEGIN';
export const FETCH_COLLECTION_SUCCESS = 'FETCH_COLLECTION_SUCCESS';
export const FETCH_COLLECTION_FAILURE = 'FETCH_COLLECTION_FAILURE';

export const fetchCollectionBegin = () => ({
  type: FETCH_COLLECTION_BEGIN
});

export const fetchCollectionSuccess = collection => ({
  type: FETCH_COLLECTION_SUCCESS,
  payload: { collection }
});

export const fetchCollectionFailure = error => ({
  type: FETCH_COLLECTION_FAILURE,
  payload: { error }
});

export const ADD_COLLECTION_BEGIN = 'ADD_COLLECTION_BEGIN';
export const ADD_COLLECTION_SUCCESS = 'ADD_COLLECTION_SUCCESS';
export const ADD_COLLECTION_FAILURE = 'ADD_COLLECTION_FAILURE';

export const addCollectionBegin = () => ({
  type: ADD_COLLECTION_BEGIN
});

export const addCollectionSuccess = collection => ({
  type: ADD_COLLECTION_SUCCESS,
  payload: { collection }
});

export const addCollectionFailure = error => ({
  type: ADD_COLLECTION_FAILURE,
  payload: { error }
});
