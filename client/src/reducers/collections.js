import {
  FETCH_COLLECTIONS_BEGIN,
  FETCH_COLLECTIONS_SUCCESS,
  FETCH_COLLECTIONS_FAILURE,
  FETCH_COLLECTION_BEGIN,
  FETCH_COLLECTION_SUCCESS,
  FETCH_COLLECTION_FAILURE
} from "../actions/collection";

const initialState = {
  all: [],
  current: null,
  loading: false,
  error: null
};

export default function collectionReducer(state = initialState, action) {
  switch (action.type) {
    case FETCH_COLLECTIONS_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_COLLECTIONS_SUCCESS:
      return {
        ...state,
        loading: false,
        all: action.payload.collections
      };

    case FETCH_COLLECTIONS_FAILURE:
      return {
        ...state,
        loading: false,
        error: action.payload.error,
        all: []
      };

    case FETCH_COLLECTION_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_COLLECTION_SUCCESS:
      return {
        ...state,
        loading: false,
        current: action.payload.collection
      };

    case FETCH_COLLECTION_FAILURE:
      return {
        ...state,
        loading: false,
        error: action.payload.error,
        current: null
      };

    default:
      return state;
  }
}
