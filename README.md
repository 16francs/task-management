# nuxt-task-management

[![CircleCI](https://circleci.com/gh/16francs/task-management.svg?style=shield)](https://circleci.com/gh/16francs/task-management)

## タスク管理システム

このシステムでは会社で作成する工程表を簡単に作成することを目的とする  

[参考サイト](https://qiita.com/GussieTech/items/23c23608daf62017230b)

----
## 機能要件   

### excelファイルへの書き出し機能
このシステムで作成した工程表はexcelで書き換えることができる  

### excelファイル読み込み機能  
excelファイルでフォーマットが合っていれば、このシステムで読み込むことができる  

### 工程表作成機能   
プロジェクトの工程表を新規作成することのできる   
ここで追加していく項目を以下に記す

- task  
なんのタスクを行うか  
- contributor  
誰が行うか  
- 期間  
どの期間の間で取り組むのか 
- 備考  
自由記述欄

### 工程管理機能  
これまでに作成したプロジェクトの工程表を管理する  
> それぞれのプロジェクトに対して、工程表は一つしか保持することができない。  
もし、ユーザーがexecelファイルをアップロードすると前のデータが書き換わる  

### 完成イメージ

![image](https://github.com/16francs/task-management/blob/master/client/app/image/image.PNG)
