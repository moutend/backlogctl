PRAGMA ENCODING="UTF-8";

CREATE TABLE IF NOT EXISTS project (
  id INTEGER NOT NULL,
  project_key STRING NOT NULL,
  name SSTRING NOT NULL,
  chart_enabled BOOLEAN NOT NULL,
  subtasking_enabled BOOLEAN NOT NULL,
  project_leader_can_edit_project_leader BOOLEAN NOT NULL,
  text_formatting_rule STRING NOT NULL,
  archived BOOLEAN NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS priority (
  id INTEGER NOT NULL,
  name STRING NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS project_status (
  id INTEGER NOT NULL,
  project_id INTEGER NOT NULL,
  name STRING NOT NULL,
  color STRING NOT NULL,
  display_order INTEGER NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS issue_type (
  id INTEGER NOT NULL,
  project_id INTEGER NOT NULL,
  name STRING NOT NULL,
  color STRING NOT NULL,
  display_order INTEGER NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user (
  id INTEGER NOT NULL,
  user_id STRING NOT NULL,
  name STRING NOT NULL,
  role_type INTEGER NOT NULL,
  language STRING NOT NULL,
  email STRING NOT NULL,
  nulab_account_id INTEGER,

  is_myself BOOLEAN NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS issue (
  id INTEGER NOT NULL,
  project_id INTEGER,
  issue_key STRING NOT NULL,
  key_id INTEGER NOT NULL,
  issue_type_id INTEGER,
  summary STRING NOT NULL,
  description STRING NOT NULL,
  resolution_id INTEGER,
  priority_id INTEGER,
  status_id INTEGER,
  assignee_id INTEGER,
  start_date DATETIME,
  due_date DATETIME,
  estimated_hours REAL,
  actual_hours REAL,
  parent_issue_id INTEGER,

  created_user_id INTEGER NOT NULL,
  updated_user_id INTEGER NOT NULL,

  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  PRIMARY KEY(id)
);
