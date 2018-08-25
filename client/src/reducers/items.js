import {
  FETCH_ITEMS_BEGIN,
  FETCH_ITEMS_SUCCESS,
  FETCH_ITEMS_FAILURE,
  FETCH_ITEM_BEGIN,
  FETCH_ITEM_SUCCESS,
  FETCH_ITEM_FAILURE
} from "../actions/item";

const initialState = {
  all: [],
  current: null,
  loading: false,
  error: null
};

export default function itemReducer(state = initialState, action) {
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

    case FETCH_ITEM_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_ITEM_SUCCESS:
      return {
        ...state,
        loading: false,
        current: action.payload.item
      };

    case FETCH_ITEM_FAILURE:
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
