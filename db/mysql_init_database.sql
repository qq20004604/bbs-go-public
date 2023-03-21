-- 用 root 权限，手动在 mysql 里执行以上代码，创建测试库和正式库
CREATE DATABASE IF NOT EXISTS bbs_test DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON bbs_test.* TO 'bbs_test_manager'@'%' IDENTIFIED BY 'fwevberbib3jbinvrewwfmievmavfv3erv';
CREATE DATABASE IF NOT EXISTS bbs DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON bbs.* TO 'bbs_manager'@'%' IDENTIFIED BY 'fwevberbib3jvavbin1e1ekwfmievmavfv3erv';