import axios from "axios";
import { ICommonGameStateWithBoard, ICoords, ISuggestion } from "./types";

const API = axios.create({
  baseURL: process.env.REACT_APP_BASE_URL || ""
});

export function makeMove(
  state: ICommonGameStateWithBoard,
  coords: ICoords
): Promise<ICommonGameStateWithBoard> {
  const cell = state.board[coords.y][coords.x];
  if (cell) return Promise.reject("cell is occupied");
  return API.post("/make-move", {
    state,
    coords
  }).then(
    ({ data }) => data,
    ({ response }) =>
      Promise.reject(response ? response.data : "failed request")
  );
}

export function suggestMoves(
  state: ICommonGameStateWithBoard,
  difficulty: number
): Promise<ISuggestion[]> {
  return API.post(`/suggest-moves?difficulty=${difficulty}`, state).then(
    ({ data }) => data,
    ({ response }) =>
      Promise.reject(response ? response.data : "failed request")
  );
}
