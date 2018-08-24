import {
  FETCH_ITEMS_BEGIN,
  FETCH_ITEMS_SUCCESS,
  FETCH_ITEMS_FAILURE
} from "../actions/item";

const initialState = {
  all: [],
  current: null,
  loading: false,
  error: null
};

export default function collectionReducer(state = initialState, action) {
  switch (action.type) {
    case FETCH_ITEMS_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_ITEMS_SUCCESS:
      return {
        ...state,
        loading: false,
        all: action.payload.items
      };

    case FETCH_ITEMS_FAILURE:
      return {
        ...state,
        loading: false,
        error: action.payload.error,
        all: []
      };

    default:
      return state;
  }
}
