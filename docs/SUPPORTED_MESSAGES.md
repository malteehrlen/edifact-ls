# Supported message specifications

Generated from the schemas registered in `internal/edifact` -- see `make docs`.
Do not hand-edit; regenerate instead.

Structural (segment/group presence, order, cardinality) validation is
available for exactly the message identities listed below -- matched on the
full (type, version, release, agency) tuple a message declares in its own
UNH. A message declaring a recognized type under a different version,
release, or agency than any row here gets an informational diagnostic
naming what IS registered, not silence and not a false match.

| Type | Version | Release | Agency | Source |
| --- | --- | --- | --- | --- |
| APERAK | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/aperak_c.htm |
| CUSCAR | D | 99B | UN | https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm |
| DELFOR | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/delfor_c.htm |
| DESADV | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/desadv_c.htm |
| IFTMCS | D | 21A | UN | https://service.unece.org/trade/untdid/d21a/trmd/iftmcs_c.htm |
| IFTMIN | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/iftmin_c.htm |
| IFTSTA | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/iftsta_c.htm |
| INVOIC | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/invoic_c.htm |
| INVRPT | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/invrpt_c.htm |
| ORDERS | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/orders_c.htm |
| ORDRSP | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/ordrsp_c.htm |
| PRICAT | D | 20A | UN | https://service.unece.org/trade/untdid/d20a/trmd/pricat_c.htm |
