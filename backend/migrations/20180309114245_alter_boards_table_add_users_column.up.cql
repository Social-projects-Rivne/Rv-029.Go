ALTER TABLE boards
    ADD (users map<uuid,text>);