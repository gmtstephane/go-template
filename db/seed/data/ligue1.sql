

INSERT INTO sport (name, type) VALUES ('Football', 'Team');

INSERT INTO championship (name, sport_id) VALUES ('Ligue 1', 1);


INSERT INTO location (name, address, city, state, country, zipcode, latitude, longitude)
VALUES
('Parc des Princes', '24 Rue du Commandant Guilbaud', 'Paris', 'Île-de-France', 'France', 75016, 48.841437, 2.253049),
('Allianz Riviera', 'Bd des Jardiniers', 'Nice', 'Provence-Alpes-Côte d''Azur', 'France', 6200, 43.704896, 7.192583),
('Stade Pierre-Mauroy', '261 Boulevard de Tournai', 'Villeneuve-d''Ascq', 'Hauts-de-France', 'France', 59650, 50.611685, 3.130215),
('Stade Louis II', '7 Avenue des Castelans', 'Monaco', 'Monaco', 'Monaco', 98000, 43.727539, 7.415535),
('Stade Francis-Le Blé', '26 Rue de Quimper', 'Brest', 'Bretagne', 'France', 29200, 48.407907, -4.495108),
('Stade Bollaert-Delelis', 'Avenue Alfred Maës', 'Lens', 'Hauts-de-France', 'France', 62300, 50.432222, 2.818889),
('Roazhon Park', '111 Rue de Lorient', 'Rennes', 'Bretagne', 'France', 35000, 48.107319, -1.712833),
('Orange Vélodrome', '3 Boulevard Michelet', 'Marseille', 'Provence-Alpes-Côte d''Azur', 'France', 13008, 43.269806, 5.396189),
('Stade Auguste-Delaune', '33 Chaussée Bocquaine', 'Reims', 'Grand Est', 'France', 51100, 49.247322, 4.024251),
('Stade de la Meinau', '12 Rue de l''Extenwoerth', 'Strasbourg', 'Grand Est', 'France', 67100, 48.559322, 7.754936),
('Groupama Stadium', '10 Avenue Simone Veil', 'Lyon', 'Auvergne-Rhône-Alpes', 'France', 69150, 45.765295, 4.982675),
('Stade Océane', 'Boulevard de Léningrad', 'Le Havre', 'Normandie', 'France', 76600, 49.512781, 0.153105),
('Stade de la Beaujoire', 'Route de Saint-Joseph', 'Nantes', 'Pays de la Loire', 'France', 44300, 47.256490, -1.524929),
('Stadium de Toulouse', '1 Allée Gabriel Biénès', 'Toulouse', 'Occitanie', 'France', 31400, 43.583322, 1.434027),
('Stade de la Mosson', '345 Avenue de Heidelberg', 'Montpellier', 'Occitanie', 'France', 34080, 43.622222, 3.811389),
('Stade du Moustoir', 'Rue du Tour des Portes', 'Lorient', 'Bretagne', 'France', 56100, 47.746667, -3.370833),
('Stade Saint-Symphorien', '3 Boulevard Saint-Symphorien', 'Metz', 'Grand Est', 'France', 57050, 49.107500, 6.158056),
('Stade Gabriel Montpied', 'Rue Adrien Mabrut', 'Clermont-Ferrand', 'Auvergne-Rhône-Alpes', 'France', 63100, 45.786111, 3.101944);


-- Insert teams with their corresponding locations for Ligue 1
INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Paris Saint-Germain', 1, id, '' FROM location WHERE name = 'Parc des Princes';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'OGC Nice', 1, id, '' FROM location WHERE name = 'Allianz Riviera';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'LOSC Lille', 1, id, '' FROM location WHERE name = 'Stade Pierre-Mauroy';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'AS Monaco', 1, id, '' FROM location WHERE name = 'Stade Louis II';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Stade Brestois 29', 1, id, '' FROM location WHERE name = 'Stade Francis-Le Blé';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'RC Lens', 1, id, '' FROM location WHERE name = 'Stade Bollaert-Delelis';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Stade Rennais FC', 1, id, '' FROM location WHERE name = 'Roazhon Park';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Olympique de Marseille', 1, id, '' FROM location WHERE name = 'Orange Vélodrome';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Stade de Reims', 1, id, '' FROM location WHERE name = 'Stade Auguste-Delaune';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'RC Strasbourg Alsace', 1, id, '' FROM location WHERE name = 'Stade de la Meinau';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Olympique Lyonnais', 1, id, '<svg version="1.1" viewBox="0 0 686.02 800" xml:space="preserve" xmlns="http://www.w3.org/2000/svg">
<g transform="matrix(.52022 0 0 .52022 .025803 0)"><path d="m0.050009 0-0.099609 1101.1c0 74.7 20.9 133.4 62.1 174.3 39.7 39.4 96.1 59.699 172.4 61.799 182.2 5.1 290.8 45.502 374.8 139.6l50.049 61 50.051-61c84-94.1 192.6-134.5 374.8-139.6 76.3-2.1 132.7-22.399 172.4-61.799 41.2-40.9 62.1-99.601 62.1-174.3l-0.0996-1101.1z" fill="#c88c28"/><path d="m32.85 32.801 0.09961 1068.3c0 65.7 17.6 116.5 52.4 151 34.5 34.3 85.099 50.501 150.2 52.301 187.7 5.3 306 46.9 398.8 151.1l24.949 30.6 24.951-30.6c92.8-104.2 211.1-145.8 398.8-151.1 65.1-1.8 115.7-18.001 150.2-52.301 34.8-34.5 52.398-85.3 52.398-151l0.1016-1068.3z" fill="#fff"/></g>
<g transform="matrix(.52022 0 0 .52022 .025803 0)"><path d="m1253.3 65.3h-1187.9l-0.1 422.5h1188z" fill="#f40043"/><g transform="translate(-.09961)" fill="#fff"><path d="m1190.1 241.1h-85.3v-123.5h84.7v21.4h-60.5v29h53.3v21.5h-53.3v30.4h61v21.2z"/><path d="m1138.5 437.1c-24.2 0-42.2-6.4-56.7-17.6l14.3-22.8c13.1 10.5 27.5 15.5 42.8 15.5 13 0 22.8-5.6 22.8-17.4s-8.3-16.8-30-22.8c-24.6-6.8-43.9-16.4-43.9-41.8s21.1-42.4 52.4-42.4c21.3 0 37.2 6 48.6 14.9l-13.1 21.5c-9.9-7.3-21.1-11.6-35.8-11.6-13.2 0-23.2 5-23.2 15.7s9.5 13.9 29.6 19.9c24.6 7.2 43.7 16.8 43.7 44.7s-20.4 44.2-51.5 44.2"/><path d="m1074.8 188c0 35.7-21.6 54.7-54.3 54.7s-54.3-19-54.3-54.7v-70.5h24.3v72.5c0 16.7 9.8 30.8 30.2 30.8s30.1-14.1 30.1-30.8v-72.4h24.1v70.4z"/><path d="m1057 435.3h-28.6v-145.6h28.6z"/><path d="m926.8 289.8-63.801 145.5h30.301l11.6-28.102h62.299l11.602 28.102h30.1l-63.301-145.5zm9.2988 40.5 21.9 52.6h-44z"/><path d="m844 436.9h-16.2l-86.7-97.3v95.6h-28v-147h16.2l86.5 97.1v-95.6h28.2z"/><path d="m649.8 117.6v123.5h24.199v-35.5l18.4 0.0996c30.6-0.2 51.5-16.098 51.5-43.398s-20.6-44.701-50.5-44.701zm24.199 21.4h18.1c16.9 0 26.701 10.601 26.701 22.801s-9.5012 22.199-26.701 22.199h-18.1z"/><path d="m675.7 436.9h-16.2l-86.7-97.3v95.6h-28v-147h16.2l86.5 97.1v-95.6h28.2z"/><path d="m440 287.9c-44.5 0-77 32.9-77 74.9s32.5 74.299 77 74.299c44.1 0 76.199-31.299 76.199-74.299s-31.699-74.9-76.199-74.9zm0 26.5c26.9 0 47 20.9 47 48.4s-19.5 48-47 48-47.801-20.5-47.801-48 20.301-48.4 47.801-48.4z"/><path d="m304.6 378.5v56.7h-28.6v-56.7l-53.4-88.8h34.1l34.8 60.6 34.8-60.6h32.3z"/><path d="m236.7 435.3h-99v-145.6h28.6v120.5h70.4z"/><path d="m194.1 116c-37.8 0-65.4 27.9-65.4 63.6s27.9 63.1 65.4 63.1 64.701-26.6 64.701-63.1-27.201-63.6-64.701-63.6zm0 22.5c22.9 0 39.9 17.7 39.9 41.1s-17 40.801-39.9 40.801c-23.4 0-40.6-17.501-40.6-40.801s17.7-41.1 40.6-41.1z"/><path d="m283 117.5h24.2v102.3h59.8v21.3h-84z"/><path d="m384.1 117.5 29.5 51.5 29.5-51.5h27.4l-45.8 75.5v48.2h-24.3v-48.2l-45.4-75.4h29.1z"/><path d="m491 117.5h15.3l46.8 81.2 47.8-81.2h15.1l3.7 123.6h-23.9l-1-74.2-34.1 59.8h-15.1l-33.4-59.8-1.2 74.2h-23.7z"/><path d="m765.2 117.5h24.2v123.6h-24.3z"/><path d="m878.9 116c-37.8 0-65.4 27.9-65.4 63.6s27.6 63.1 65.4 63.1c5.6 0 10.9-0.49922 16-1.6992l31 28.1 15-17.4-23-20.898c16-11.1 25.699-29.601 25.699-51.201 0-36.5-27.199-63.6-64.699-63.6zm0.09961 22.5c22.9 0 39.9 17.9 39.9 41.1s-17 40.801-39.9 40.801c-23.4 0-40.6-17.601-40.6-40.801s17.7-41.1 40.6-41.1z"/></g><path d="m65.35 520.3v580.8c0 114.3 57.4 166.6 170 170.1 196.1 6.3 330.35 50.701 423.95 158.3 93.6-107.6 227.85-152 423.95-158.3 112.6-3.5 170-55.8 170-170.1v-580.8z" fill="#0f23aa"/><path d="m1143.1 1160.8h-332.1v-550.6h160.3v411.1h171.8z" fill="#fff"/><path d="m442.5 1173.9c-159 0-288.4-129.4-288.4-288.4s129.4-288.4 288.4-288.4 288.4 129.4 288.4 288.4-129.4 288.4-288.4 288.4" fill="#fff"/><g fill="#c88c28"><path class="st1" d="m384.9 770.5c11.9-3.6 21.9-8.7 33.6-4.8-4.3-8.7-14.4-12.7-23.6-10.1-4.8 1.3-9.1 3.5-13 4.8-11.9 5.3-23.1 5.3-35.3 1.7 9 11.1 25.7 12.7 38.3 8.4"/><path class="st1" d="m628.5 806.3c-16.6 33.2-82.1 27.1-78.2 65l-4.8-8.7c-5.6-10.5-1.3-23.6 9.2-29.2l60.2-31.8c19.2-10.1 26.2-35.4 13.5-54.2-10.1-14.8-29.3-20-45.9-13.9 0 0-17.5 7.8-24.9-12.7-8.7 9.2-4.3 21.8 1.3 25.4-5.6 0.8-16.6 8.6-17 24 12.7 3.5 26.6-1.3 29.2-5.3-0.8 4.8 7.4 14.8 19.3 12.7-10.1-10.5-7.4-22.3 0-25.3 6.9-2.6 15.7-0.5 20.5 5.6 6.6 8.7 3.5 21-5.8 25.7l-60.2 31.8c-20.5 10.5-28.4 36.3-17.4 56.8l27.1 51.1c3 5.6 2.2 13.1-2.6 17.5-6.1 6.5-16.6 5.2-21.9-2.3-16.5-24-30.5-49.7-41.4-77.2 9.2 3.5 16.2-2.6 16.2-2.6s-10.5-7.4-4.8-22.6c3-9.2 4-17.1 2.6-24.5 2.6 2.3 6.6 4 11.4 3.6-4-18 0.8-18.8 2.2-34.5 0.3-7.1-1.3-12.3-3.5-16.2 3.8 0.4 8.3 0 12.2-4-17.1-7.1-14-31-34.5-33.7 7.8-3.5 8.3-10.4 8.3-11.7-17.5 1.3-30.2-11.4-49-11.4-20.5 0-26.1 5.3-33.5 5.3-10.5 0-15.7 5.2-14.4 10.4l-17.5 4.8c-2.6 4.8 0 8.7 2.2 10.1 0 0-3 2.6-1.8 7 1.3 4.9 4.4 13.5 4.4 13.5l3.5-3.5c1.8-1.7 4.3-3 7.4-3.5l5.6-0.5c7.1-0.8 13.5 2.6 17.1 9.2l7.4 14.8h-37.6s2.6 8.3 3.5 10.9c1.8 4 5.3 5.8 11.9 5.8 7 0 10.9-1 10.9-1v10.1c0 5.6-4.4 13-19.6 8.7l-31.9-9.2c-18.8-4.8-28-13.5-40.7-29.2l-3.8-6.1c-3.1-4-4.4-8.7-4-13.5 0.4-4 0.4-9.1-0.5-14-1.7-11.7-10.4-15.7-12.7-16.2 0.5 1.3 4.4 11.9 0 21.4-3.5-10.4-10.4-18.3-26.6-19.6 0 0 11.9 10.1 12.2 23.1-14.4-9.6-31.9-3-31.9-3s27.9 5.6 27.9 27.9c0 8.3-4.3 19.3-17.4 26.7 0 0 13.9 3 23.1-2.3 2.6-1.7 6.1-1.7 8.7 0.5 9.1 9.2 29.2 28.9 32.2 35 6.6 13.1-2.2 23.6-2.2 23.6s14-1 18.8-9.2c0.8-1.3 2.6-1.8 4-1.3l26.1 15.7-17.6 11.8c-15.7 10.4-34 15.7-52.5 15.7h-17.8c-5.6 0-11.4-1.3-16.2-4-4.8-2.2-12.2-5.3-19.6-5.3-15.7 0-20.5 7.9-20.5 7.9s15.7-0.8 21.8 5.3c-9.6-1.3-27.1 3-28.9 18.8 0 0 11.4-6.6 22.3-4.9-7.8 3.6-16.5 14-6.1 30.7 0 0 1-17.5 22.8-17.5s18.8 19.2 18.8 19.2 9.6-2.2 11.7-14.4c0.5-2.6 2.2-4.8 4.8-4.8h45.9c1.8 0 3 1.3 3 2.6 2.3 9.6 15 12.2 15 12.2s-6.2-11.4 3.5-15.7l33.5-14.5 15.3 20.1c-25.3 14.4-61.1 30.5-70.8 34.5-20 7.9-25.8 7.9-34.5 7.4 0 0 10.5 9.6 30.2 7.4 1.8 0 3.5-0.5 5.6-1.3 1.8-0.8 4 0 4.8 1.8 3.1 6.1 7.9 15.7 9.2 19.6 3.1 8.3-2.6 10.4-8.7 13.1-8.3 4-17 7.9-31 12.7-21.3-12.2-37.5-0.8-37.5-0.8s12.6 2.2 18.7 9.6c-12.2-1-25.3 6.1-29.7 20 3.1-1.3 16.6-6.9 28.4-5.2-3.5 3-8.3 9.6-8.3 18.3 0.5 4.3 1.3 6.5 1.3 6.5s10.5-16.5 29.7-16.5c8.3 0 12.2 4.3 12.2 8.7 0 3.5-2.6 7.8-2.6 7.8s25.4-2.2 17.1-24.8l46.2-16.6c2.2-0.8 4.8-0.5 6.1 1.3 3.1 2.6 8.7 6.1 16.6 5.3l-13.5-30.5c-3.6-7.9 0.8-16.6 9.9-19.3l24.1-7.8s6.5 5.2 27.5 22.6c13.9 11.4 13.9 31.5 13.9 31.5s12.2-4.8 12.2-15.8c2.3 1 4.9 1.8 7.4 2.3 12.2 3 20.6 13.5 12.2 30.5l-7.8 14.4c-3.1 5.3-6.1 6.6-10.5 6.6-4.8-0.5-10.9-0.8-17 0-22.6 2.2-23.1 17.5-23.1 17.5s7.4-4.4 16.5-4.4c3.6 0 7.1 1 10.1 1.8 0 0-18.8 3-18.8 19.2 0 6.6 5.3 10.1 5.3 10.1s4.8-17 24.1-17c0 0-10.5 4.8-10.5 14.8s9.2 13.1 9.2 13.1-6.1-15.2 23.1-15.2h7.9c5.2 0 9.6-3.1 11.4-7.9l5.6-15.7c3.5-9.2 9.2-9.2 12.7-8.4 5.6 1.3 5.6 9.2 5.6 9.2s9.6-14.4-1.8-27.5c-2.2-2.6-3-6.9-1.7-10.4l7.9-21.8c17.4 4.3 25.8-10.1 25.8-10.1-18.5-4-27.5-9.7-54.1-31.5-11.8-9.6-19.3-21.8-19.3-24 6.6 7 15.7 10.9 25.3 10.9 7.4 0 14.8-2.6 21.4-7.4 5.8-4.8 9.6-10.9 11.4-17.5 9.2-15.7 31-3.1 34.5-35.3-9.2 13-29.2 4.3-33.2 20.5-8.3-35.8 43.2-12.2 43.7-66.3-12.7 21.8-57.7 14.4-54.6 41.4l-0.5-0.4c-17.4-56.1 68.2-19.9 63.4-94m-211.8-76.1c-6.5 0-8.3-9.6 1.8-9.6 7.1 0 13.5 2.2 13.5 2.2s-9.1 7.4-15.3 7.4"/></g></g></svg>' FROM location WHERE name = 'Groupama Stadium';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Le Havre AC', 1, id, '' FROM location WHERE name = 'Stade Océane';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'FC Nantes', 1, id, '' FROM location WHERE name = 'Stade de la Beaujoire';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Toulouse FC', 1, id, '' FROM location WHERE name = 'Stadium de Toulouse';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Montpellier HSC', 1, id, '' FROM location WHERE name = 'Stade de la Mosson';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'FC Lorient', 1, id, '' FROM location WHERE name = 'Stade du Moustoir';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'FC Metz', 1, id, '' FROM location WHERE name = 'Stade Saint-Symphorien';

INSERT INTO team (name, sport_id, location_id, icon_svg)
SELECT 'Clermont Foot 63', 1, id, '' FROM location WHERE name = 'Stade Gabriel Montpied';
