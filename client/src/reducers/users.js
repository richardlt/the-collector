import {
  FETCH_ME_BEGIN,
  FETCH_ME_SUCCESS,
  FETCH_ME_FAILURE
} from '../actions/user';

const initialState = {
  me: null,
  loading: false,
  error: null
};

export default function userReducer(state = initialState, action) {
  switch (action.type) {
    case FETCH_ME_BEGIN:
      return {
        ...state,
        loading: true,
        error: null
      };

    case FETCH_ME_SUCCESS:
      return {
        ...state,
        loading: false,
        me: action.payload.me
      };

    case FETCH_ME_FAILURE:
      return {
        ...state,
        loading: false,
        error: action.payload.error,
        me: null
      };

    default:
      return state;
  }
}
