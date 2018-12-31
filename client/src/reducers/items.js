import {
  FETCH_ITEMS_BEGIN,
  FETCH_ITEMS_SUCCESS,
  FETCH_ITEMS_FAILURE,
  FETCH_ITEM_BEGIN,
  FETCH_ITEM_SUCCESS,
  FETCH_ITEM_FAILURE,
  ADD_ITEM_BEGIN,
  ADD_ITEM_SUCCESS,
  ADD_ITEM_FAILURE,
  UPDATE_ITEM_FILE_BEGIN,
  UPDATE_ITEM_FILE_SUCCESS,
  UPDATE_ITEM_FILE_FAILURE,
  DELETE_ITEM_BEGIN,
  DELETE_ITEM_SUCCESS,
  DELETE_ITEM_FAILURE
} from '../actions/item';

const initialState = {
  all: [],
  loading: false,
  error: null
};

export default function itemReducer(state = initialState, action) {
  switch (action.type) {
    case FETCH_ITEMS_BEGIN:
    case FETCH_ITEM_BEGIN:
    case ADD_ITEM_BEGIN:
    case UPDATE_ITEM_FILE_BEGIN:
    case DELETE_ITEM_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_ITEMS_FAILURE:
    case FETCH_ITEM_FAILURE:
    case ADD_ITEM_FAILURE:
    case UPDATE_ITEM_FILE_FAILURE:
    case DELETE_ITEM_FAILURE:
      return {
        ...state,
        loading: false,
        error: action.payload.error
      };

    case FETCH_ITEMS_SUCCESS:
      return {
        ...state,
        loading: false,
        all: action.payload.items
      };

    case FETCH_ITEM_SUCCESS:
    case UPDATE_ITEM_FILE_SUCCESS:
      return {
        ...state,
        loading: false,
        all: state.all
          .filter(i => i.uuid != action.payload.item.uuid)
          .concat(action.payload.item)
      };

    case ADD_ITEM_SUCCESS:
      return {
        ...state,
        loading: false,
        all: state.all.concat(action.payload.item)
      };

    case DELETE_ITEM_SUCCESS:
      return {
        ...state,
        loading: false,
        all: state.all.filter(i => i.uuid != action.payload.uuid)
      };

    default:
      return state;
  }
}
