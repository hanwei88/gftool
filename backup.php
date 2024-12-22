<?php

$path = dirname(__FILE__);
var_dump($path);

$file = fopen($path  . '/backup.txt', 'r');
if (!$file) {
	exit('文件不存在');
}



// 备份目录
$backpath = $path . '/../gitbackup/' . date('YmdHis');
if (!file_exists($backpath)) {
	mkdir($backpath, true, 0755);
}
$bakfiles = [];

while (!feof($file)) {
	$line = trim(fgets($file));
	if (!$line  ||  !file_exists($path . '/' . $line)) {
		continue;
	}
	// 创建目标文件目录
	if (strpos($line, '/') !== false && !file_exists($backpath . '/' . dirname($line))) {
		mkdir($backpath . '/' . dirname($line), true, 0755);
	}
	if (copy($path . '/' . $line, $backpath . '/' . $line)) {
		$bakfiles[] = $line;
		unlink($path . '/' . $line);
	}
}

fclose($file);

echo "备份目录: {$backpath}\n";
echo "备份文件列表:\n";
var_dump($bakfiles);