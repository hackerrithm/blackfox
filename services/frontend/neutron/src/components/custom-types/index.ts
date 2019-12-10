export type UserAuth = {
    id: number;
    user: string;
  };
  
  export type ToDoItemType = { id: string; toDo?: string; complete?: boolean };
  
  export enum ActionType {
    add = "ADD",
    delete = "DELETE",
    updateStatus = "UPDATE"
  }