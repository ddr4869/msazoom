
const CreateBoardForm = ({ handleSubmitBoardForm, setShowCreateBoardForm }) => (
  <form onSubmit={handleSubmitBoardForm}>
    <input type="text" placeholder="Board Name" name="board_name" required />
    <input type="password" placeholder="Board Password" name="board_password" required />
    <button type="submit">Create</button>
    <button type="button" onClick={() => setShowCreateBoardForm(false)}>Cancel</button>
  </form>
);

export default CreateBoardForm;
