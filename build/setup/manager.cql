INSERT INTO scyna2.application (code, auth_url) VALUES ('scyna_manager', '/scyna2/manager/login');

INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/domain/create');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/create');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/update');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/change-secret');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/delete');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/list');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/get');
INSERT INTO scyna2.client_use_endpoint (client, url) VALUES ('admin', '/scyna2/manager/module/update-config');
