ALTER TABLE issues
    ADD (logs map<uuid,text>);
