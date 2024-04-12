-- Project Name : 技術書典16サンプル
-- Date/Time    : 2024/04/12 20:29:06
-- Author       : 合同会社モイテリカ
-- RDBMS Type   : PostgreSQL
-- Application  : A5:SQL Mk-2

create table "Enumログ区分" (
  "区分" char(1) not null
  , "名称" text not null
  , constraint "Enumログ区分_PKC" primary key ("区分")
) ;

create table "Enum受払区分" (
  "受払区分" integer not null
  , "名称" text not null
  , "受入フラグ" boolean default false not null
  , "符号" integer not null check ("符号" in (-1,0,1))
  , constraint "Enum受払区分_PKC" primary key ("受払区分")
) ;

create table "gorp_migrations" (
  "id" text not null
  , "applied_at" timestamp(6) with time zone
  , constraint "gorp_migrations_PKC" primary key ("id")
) ;

create table "リソース変更履歴" (
  "No" bigserial not null
  , "登録日時" timestamp with time zone default now() not null
  , "リソース名" text not null
  , "変更区分" char(1) not null
  , "変更内容" jsonb not null
  , constraint "リソース変更履歴_PKC" primary key ("No")
) ;

create table "リソース変更履歴_アップロード時" (
  "No" bigint not null
  , "アップロード履歴ID" bigint not null
  , constraint "リソース変更履歴_アップロード時_PKC" primary key ("No")
) ;

create table "ログ" (
  "No" bigserial not null
  , "登録日時" timestamp with time zone default now() not null
  , "区分" char(1) default 'I' not null
  , "内容" text not null
  , constraint "ログ_PKC" primary key ("No")
) ;

create table "ログ_アップロード時" (
  "No" bigint not null
  , "アップロード履歴ID" bigint not null
  , constraint "ログ_アップロード時_PKC" primary key ("No")
) ;

create table "ログ_画面操作時" (
  "No" bigint not null
  , "操作ユーザーID" bigint not null
  , constraint "ログ_画面操作時_PKC" primary key ("No")
) ;

create table "単位" (
  "ID" bigserial not null
  , "コード" text not null
  , "名称" text not null
  , constraint "単位_PKC" primary key ("ID")
) ;

create unique index "単位_IX1"
  on "単位"("コード");

create table "受払" (
  "No" bigserial not null
  , "登録日時" timestamp with time zone default now() not null
  , "計上月" timestamp with time zone not null
  , "受払区分" integer not null
  , "赤伝フラグ" boolean default false not null
  , "品目ID" bigint not null
  , "基準数量" numeric(20,6) not null
  , "基準単位ID" smallint not null
  , constraint "受払_PKC" primary key ("No")
) ;

create table "受払_仕入" (
  "No" bigint not null
  , "仕入数量" numeric(20,6) not null
  , "仕入単位ID" smallint not null
  , "仕入金額" bigint default 0 not null
  , "仕入通貨ID" smallint not null
  , "仕入単価" numeric(20,6) default 0 not null
  , constraint "受払_仕入_PKC" primary key ("No")
) ;

create table "受払_出荷" (
  "No" bigint not null
  , "出荷数量" numeric(20,6) not null
  , "出荷単位ID" smallint not null
  , constraint "受払_出荷_PKC" primary key ("No")
) ;

create table "受払_投入実績" (
  "No" bigint not null
  , "投入数量" numeric(20,6) not null
  , "投入単位ID" smallint not null
  , "製造指図ID" bigint not null
  , constraint "受払_投入実績_PKC" primary key ("No")
) ;

create table "受払_製造実績" (
  "No" bigint not null
  , "製造数量" numeric(20,6) not null
  , "製造単位ID" smallint not null
  , "製造指図ID" bigint not null
  , constraint "受払_製造実績_PKC" primary key ("No")
) ;

create table "品目" (
  "ID" bigserial not null
  , "コード" text not null
  , "名称" text not null
  , "基準単位ID" smallint not null
  , "生産用品目区分ID" bigint not null
  , constraint "品目_PKC" primary key ("ID")
) ;

create unique index "品目_IX1"
  on "品目"("コード");

create table "品目_仕入品" (
  "ID" bigint not null
  , "標準単価" numeric(20,6) not null
  , "標準単価通貨ID" smallint not null
  , "標準単価単位ID" smallint not null
  , constraint "品目_仕入品_PKC" primary key ("ID")
) ;

create table "品目_製造品" (
  "ID" bigint not null
  , "MRP計算対象フラグ" boolean default false not null
  , constraint "品目_製造品_PKC" primary key ("ID")
) ;

create table "生産用品目区分" (
  "ID" bigserial not null
  , "コード" text not null
  , "名称" text not null
  , "何かのフラグ1" boolean default false not null
  , "何かのフラグ2" boolean default false not null
  , constraint "生産用品目区分_PKC" primary key ("ID")
) ;

create unique index "生産用品目区分_IX1"
  on "生産用品目区分"("コード");

comment on column "Enumログ区分"."区分" is '';
comment on column "Enumログ区分"."名称" is '';

comment on column "Enum受払区分"."受払区分" is '';
comment on column "Enum受払区分"."名称" is '';
comment on column "Enum受払区分"."受入フラグ" is ':true⇒受入、false⇒払出';
comment on column "Enum受払区分"."符号" is '';

comment on column "gorp_migrations"."id" is '';
comment on column "gorp_migrations"."applied_at" is '';

comment on column "リソース変更履歴"."No" is '';
comment on column "リソース変更履歴"."登録日時" is '';
comment on column "リソース変更履歴"."リソース名" is '';
comment on column "リソース変更履歴"."変更区分" is ':A, M , D';
comment on column "リソース変更履歴"."変更内容" is '';

comment on column "リソース変更履歴_アップロード時"."No" is '';
comment on column "リソース変更履歴_アップロード時"."アップロード履歴ID" is '';

comment on column "ログ"."No" is '';
comment on column "ログ"."登録日時" is '';
comment on column "ログ"."区分" is '';
comment on column "ログ"."内容" is '';

comment on column "ログ_アップロード時"."No" is '';
comment on column "ログ_アップロード時"."アップロード履歴ID" is '';

comment on column "ログ_画面操作時"."No" is '';
comment on column "ログ_画面操作時"."操作ユーザーID" is '';

comment on column "単位"."ID" is '';
comment on column "単位"."コード" is '';
comment on column "単位"."名称" is '';

comment on column "受払"."No" is '';
comment on column "受払"."登録日時" is '';
comment on column "受払"."計上月" is '';
comment on column "受払"."受払区分" is '';
comment on column "受払"."赤伝フラグ" is '';
comment on column "受払"."品目ID" is '';
comment on column "受払"."基準数量" is '';
comment on column "受払"."基準単位ID" is '';

comment on column "受払_仕入"."No" is '';
comment on column "受払_仕入"."仕入数量" is '';
comment on column "受払_仕入"."仕入単位ID" is '';
comment on column "受払_仕入"."仕入金額" is '';
comment on column "受払_仕入"."仕入通貨ID" is '';
comment on column "受払_仕入"."仕入単価" is '';

comment on column "受払_出荷"."No" is '';
comment on column "受払_出荷"."出荷数量" is '';
comment on column "受払_出荷"."出荷単位ID" is '';

comment on column "受払_投入実績"."No" is '';
comment on column "受払_投入実績"."投入数量" is '';
comment on column "受払_投入実績"."投入単位ID" is '';
comment on column "受払_投入実績"."製造指図ID" is '';

comment on column "受払_製造実績"."No" is '';
comment on column "受払_製造実績"."製造数量" is '';
comment on column "受払_製造実績"."製造単位ID" is '';
comment on column "受払_製造実績"."製造指図ID" is '';

comment on column "品目"."ID" is '';
comment on column "品目"."コード" is '';
comment on column "品目"."名称" is '';
comment on column "品目"."基準単位ID" is '';
comment on column "品目"."生産用品目区分ID" is '';

comment on column "品目_仕入品"."ID" is '';
comment on column "品目_仕入品"."標準単価" is ':発注単位あたりの単価';
comment on column "品目_仕入品"."標準単価通貨ID" is '';
comment on column "品目_仕入品"."標準単価単位ID" is '';

comment on column "品目_製造品"."ID" is '';
comment on column "品目_製造品"."MRP計算対象フラグ" is '';

comment on column "生産用品目区分"."ID" is '';
comment on column "生産用品目区分"."コード" is '';
comment on column "生産用品目区分"."名称" is '';
comment on column "生産用品目区分"."何かのフラグ1" is '';
comment on column "生産用品目区分"."何かのフラグ2" is '';

