DROP TABLE if exists public.report;
CREATE TABLE IF NOT EXISTS public.report (
  id                  SERIAL4   NOT NULL,
  run_id              BIGINT    NOT NULL,
  group_id            TEXT      NOT NULL,
  title               TEXT      NOT NULL,
  description         VARCHAR,
  control_id          TEXT      NOT NULL,
  control_title       TEXT      NOT NULL,
  control_description TEXT      NOT NULL,
  reason              TEXT      NOT NULL,
  resource            TEXT      NOT NULL,
  status              TEXT      NOT NULL,
  severity            VARCHAR,
  compartment         VARCHAR,
  name                VARCHAR,
  region              VARCHAR,
  reporting_region    VARCHAR,
  tenant              TEXT      NOT NULL,
  category            TEXT      NOT NULL,
  cis                 BOOLEAN   NOT NULL,
  cis_item_id         DECIMAL   NOT NULL,
  cis_level           BIGINT    NOT NULL,
  cis_section_id      BIGINT    NOT NULL,
  cis_type            TEXT      NOT NULL,
  cis_version         TEXT      NOT NULL,
  plugin              TEXT      NOT NULL,
  service             TEXT      NOT NULL,
  type                VARCHAR,
  created             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_report PRIMARY KEY (id)
);

DROP TABLE if EXISTS public.run_id;
CREATE TABLE IF NOT EXISTS public.run_id (
  id BIGINT NOT NULL
);

INSERT INTO public.run_id ( id ) VALUES ( 0 );